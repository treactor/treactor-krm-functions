// This file will be processed and embedded to pluginator.

package main

import (
	"fmt"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"github.com/go-errors/errors"
	"github.com/treactor/treactor-kpt-functions/common"
	"github.com/treactor/treactor-kpt-functions/common/constants"
	"github.com/treactor/treactor-kpt-functions/create-atoms/element"
	"os"
	gwv1beta1 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	"strconv"
	"strings"
)

const (
	legacyFnConfigKind = "SetAnnotationConfig"
	fnConfigKind       = "SetAnnotations"
	fnUri              = constants.FnUriPrefix + "create-atoms"
)

func upsert(env []map[string]interface{}, envVar map[string]interface{}) []map[string]interface{} {
	name, found := envVar["name"]
	if !found {
		return env
	}
	for ix, e := range env {
		if name == e["name"] {
			env[ix] = envVar
			return env
		}
	}
	return append(env, envVar)
}

type CreateAtomsFn struct {
	config *CreateAtomsConfig
	items  []*fn.KubeObject
}

func CreateFn(config *CreateAtomsConfig) *CreateAtomsFn {
	function := CreateAtomsFn{
		items:  []*fn.KubeObject{},
		config: config,
	}
	return &function
}

func (function *CreateAtomsFn) ensureServiceForElement(rl *fn.ResourceList, atom element.Element) error {
	app := function.getAppName(atom)

	service, found := common.FindObject(rl, "", "v1", "Service", app)
	if !found {
		var err error
		service, err = fn.ParseKubeObject([]byte(`# Generate by gcr.io/treactor/kpt-fn/create-atoms
apiVersion: v1
kind: Service
metadata:
  labels:
    app: atom
  name: atom
  namespace: generated
spec:
  ports:
    - name: http
      port: 80
      protocol: TCP
      targetPort: 30000
      appProtocol: http
  selector:
    app: atom
`))
		if err != nil {
			return err
		}
	}
	service.SetAnnotation("config.kubernetes.io/managed-by", fnUri)
	service.SetName(app)
	service.SetLabel("app", app)
	service.SetNestedField(app, "spec", "selector", "app")
	function.items = append(function.items, service)
	return nil
}

func (function *CreateAtomsFn) ensureRouteForElement(rl *fn.ResourceList, atom element.Element) error {
	app := function.getAppName(atom)

	route, found := common.FindObject(rl, "gateway.networking.k8s.io", "v1alpha2", "HTTPRoute", app)
	if !found {
		var err error
		route, err = fn.ParseKubeObject([]byte(`# Generate by gcr.io/treactor/kpt-fn/create-atoms
apiVersion: gateway.networking.k8s.io/v1alpha2
kind: HTTPRoute
metadata:
  name: atom
  namespace: generated
spec:
  parentRefs:
    - name: gateway
  hostnames: ["treactor.example.com"]
`))
		if err != nil {
			return err
		}
		route.SetName(app)
	}

	route.SetAnnotation("config.kubernetes.io/managed-by", fnUri)

	path := fmt.Sprintf("/treact/nodes/%d/info", atom.Number)
	port := gwv1beta1.PortNumber(80)

	var rules []gwv1beta1.HTTPRouteRule
	rules = append(rules, gwv1beta1.HTTPRouteRule{
		Matches: []gwv1beta1.HTTPRouteMatch{{
			Path: &gwv1beta1.HTTPPathMatch{
				Type:  nil,
				Value: &path,
			},
		},
		},
		BackendRefs: []gwv1beta1.HTTPBackendRef{{
			BackendRef: gwv1beta1.BackendRef{
				BackendObjectReference: gwv1beta1.BackendObjectReference{
					Name: gwv1beta1.ObjectName(app),
					Port: &port,
				},
			},
		},
		},
	})
	if atom.Number == 501 {
		pathInfo := "/treact/info"
		rules[0].Matches = append(rules[0].Matches, gwv1beta1.HTTPRouteMatch{
			Path: &gwv1beta1.HTTPPathMatch{
				Type:  nil,
				Value: &pathInfo,
			},
		})
	}
	route.SetNestedField(rules, "spec", "rules")

	function.items = append(function.items, route)
	return nil
}

func (function *CreateAtomsFn) getAppName(element element.Element) string {
	if element.Number < 200 {
		return fmt.Sprintf("atom-%s", strings.ToLower(element.Symbol))
	} else if element.Number >= 200 && element.Number < 299 {
		return fmt.Sprintf("bond-%d", element.Number-200)
	} else if element.Number == 299 {
		return "bond-n"
	} else if element.Number == 500 {
		return "treactor-ui"
	} else if element.Number == 501 {
		return "treactor-api"
	}
	return "unknown"
}

func (function *CreateAtomsFn) ensureDeploymentForElement(rl *fn.ResourceList, atom element.Element) error {
	app := function.getAppName(atom)

	deployment, found := common.FindObject(rl, "app", "v1", "Deployment", app)
	if !found {
		var err error
		deployment, err = fn.ParseKubeObject([]byte(`# Generate by gcr.io/treactor/kpt-fn/create-atoms
apiVersion: apps/v1
kind: Deployment
metadata:
  name: atom
  namespace: generated
spec:
  template:
    spec:
      containers:
        - name: atom
          image: treactor/treactor
          env:
            - name: PORT
              value: '30000'
            - name: SERVICE_VERSION
              value: '0.0.0'
            - name: TREACTOR_MODE
              value: 'cluster'
            - name: TREACTOR_TRACE_PROPAGATION
              value: 'w3c'
            - name: TREACTOR_TRACE_INTERNAL
              value: '1'
            - name: TREACTOR_RANDOM
              value: 'nothing_secret'
          ports:
            - containerPort: 30000
              name: http
              protocol: TCP
          readinessProbe:
            httpGet:
              path: /healthz
              port: http
`))
		if err != nil {
			return err
		}
		deployment.SetName(app)
	}

	deployment.SetAnnotation("config.kubernetes.io/managed-by", fnUri)
	deployment.SetNestedField(app, "spec", "selector", "matchLabels", "app")
	deployment.SetNestedField(app, "spec", "template", "metadata", "labels", "app")
	containers, _, err := deployment.NestedSlice("spec", "template", "spec", "containers")
	for _, container := range containers {
		if container.GetString("name") == "atom" {

			container.SetNestedField(function.config.ImageForAtom(atom), "image")
			var env []map[string]interface{}
			_, err = container.Get(&env, "env")
			if err != nil {
				return err
			}
			env = upsert(env, map[string]interface{}{
				"name":  "SERVICE_NAME",
				"value": app,
			})
			env = upsert(env, map[string]interface{}{
				"name":  "SERVICE_VERSION",
				"value": function.config.ImageTagForAtom(atom),
			})
			env = upsert(env, map[string]interface{}{
				"name":  "TREACTOR_NUMBER",
				"value": strconv.Itoa(int(atom.Number)),
			})
			env = upsert(env, map[string]interface{}{
				"name":  "TREACTOR_MAX_NUMBER",
				"value": strconv.Itoa(int(function.config.MaxNumber)),
			})
			env = upsert(env, map[string]interface{}{
				"name":  "TREACTOR_MAX_BOND",
				"value": strconv.Itoa(int(function.config.MaxBond)),
			})
			env = upsert(env, map[string]interface{}{
				"name":  "TREACTOR_COMPONENT",
				"value": strings.ToLower(atom.Symbol),
			})
			module := "atom"
			if atom.Module != "" {
				module = atom.Module
			}
			env = upsert(env, map[string]interface{}{
				"name":  "TREACTOR_MODULE",
				"value": module,
			})

			container.SetNestedField(env, "env")
			if err != nil {
				return err
			}
			function.items = append(function.items, deployment)
			return nil
		}
	}
	return errors.New("No container atom found")
}

func Run(rl *fn.ResourceList) (bool, error) {
	var config CreateAtomsConfig
	err, _ := config.Config(rl.FunctionConfig)
	if err != nil {
		return false, err
	}
	elements := element.ReadElements()
	var function = CreateFn(&config)
	for _, element := range elements.Elements {
		if element.Number <= config.MaxNumber {
			err := function.ensureDeploymentForElement(rl, element)
			if err != nil {
				return false, err
			}
			err = function.ensureServiceForElement(rl, element)
			if err != nil {
				return false, err
			}
			err = function.ensureRouteForElement(rl, element)
			if err != nil {
				return false, err
			}
		}
		if element.Number >= 200 && element.Number < 300 {
			if element.Number <= 200+config.MaxBond || element.Number == 299 {
				err := function.ensureDeploymentForElement(rl, element)
				if err != nil {
					return false, err
				}
				err = function.ensureServiceForElement(rl, element)
				if err != nil {
					return false, err
				}
				err = function.ensureRouteForElement(rl, element)
				if err != nil {
					return false, err
				}
			}
		}
		if element.Number >= 500 && element.Number < 600 {
			err := function.ensureDeploymentForElement(rl, element)
			if err != nil {
				return false, err
			}
			err = function.ensureServiceForElement(rl, element)
			if err != nil {
				return false, err
			}
			if element.Number != 500 {
				err = function.ensureRouteForElement(rl, element)
				if err != nil {
					return false, err
				}
			}
		}
	}

	var items []*fn.KubeObject
	for _, item := range rl.Items {
		if item.GetAnnotation("config.kubernetes.io/managed-by") != fnUri {
			items = append(items, item)
		}
	}
	rl.Items = append(items, function.items...)

	return true, nil
}

func main() {
	if err := fn.AsMain(fn.ResourceListProcessorFunc(Run)); err != nil {
		os.Exit(1)
	}
}
