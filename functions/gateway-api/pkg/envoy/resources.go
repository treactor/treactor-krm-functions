package envoy

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"github.com/treactor/treactor-kpt-functions/common"
	"github.com/treactor/treactor-kpt-functions/gateway-api/pkg/config"
	"github.com/treactor/treactor-kpt-functions/gateway-api/pkg/fnc"
	"github.com/treactor/treactor-kpt-functions/gateway-api/pkg/routes"
)

type resources struct {
	items []*fn.KubeObject
}

func Create(rl *fn.ResourceList, c *config.EnvoyConfig, routes routes.Routes) ([]*fn.KubeObject, error) {
	r := resources{items: []*fn.KubeObject{}}

	err := r.ensureConfigMapForEnvoy(rl, routes)
	if err != nil {
		return nil, err
	}
	err = r.ensureDeploymentForEnvoy(rl)
	if err != nil {
		return nil, err
	}
	err = r.ensureServiceForEnvoy(rl)
	if err != nil {
		return nil, err
	}
	if c != nil && c.EnableIngress {
		err = r.ensureIngressForEnvoy(rl)
	}
	if err != nil {
		return nil, err
	}

	return r.items, nil
}

func (r *resources) ensureServiceForEnvoy(rl *fn.ResourceList) error {
	kubeObject, found := common.FindObject(rl, "", "v1", "Service", "envoy")
	if !found {
		var err error
		kubeObject, err = fn.ParseKubeObject([]byte(`# Generate by gcr.io/treactor/kpt-fn/gateway-api
apiVersion: v1
kind: Service
metadata:
  labels:
    app: envoy
  name: envoy
  namespace: generated
spec:
  ports:
    - name: http
      port: 30000
      protocol: TCP
      targetPort: 30000
  selector:
    app: envoy
`))
		if err != nil {
			return err
		}
		kubeObject.SetAnnotation("config.kubernetes.io/managed-by", fnc.FnUri)
	}
	r.items = append(r.items, kubeObject)
	return nil
}

func (r *resources) ensureDeploymentForEnvoy(rl *fn.ResourceList) error {
	kubeObject, found := common.FindObject(rl, "apps", "v1", "Deployment", "envoy")
	if !found {
		var err error
		kubeObject, err = fn.ParseKubeObject([]byte(`# Generate by gcr.io/treactor/kpt-fn/gateway-api
apiVersion: apps/v1
kind: Deployment
metadata:
  name: envoy
  namespace: generated
spec:
  template:
    spec:
      containers:
      - name: envoy
        image: envoyproxy/envoy-distroless:v1.23-latest
        env:
        - name: PORT
          value: "30000"
        ports:
        - name: http
          containerPort: 30000
          protocol: TCP
        - name: admin
          containerPort: 9901
          protocol: TCP
        readinessProbe:
          httpGet:
            path: /ready
            port: admin
        volumeMounts:
        - name: config-volume
          mountPath: /etc/envoy
      volumes:
      - name: config-volume
        configMap:
          name: envoy
    metadata:
      labels:
        app: envoy
  selector:
    matchLabels:
      app: envoy
`))
		if err != nil {
			return err
		}
		kubeObject.SetAnnotation("config.kubernetes.io/managed-by", fnc.FnUri)
	}
	r.items = append(r.items, kubeObject)
	return nil
}

func (r *resources) ensureIngressForEnvoy(rl *fn.ResourceList) error {
	kubeObject, found := common.FindObject(rl, "networking.k8s.io", "v1", "Ingress", "treactor")
	if !found {
		var err error
		kubeObject, err = fn.ParseKubeObject([]byte(`# Generate by gcr.io/treactor/kpt-fn/gateway-api
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: treactor
  namespace: generated
spec:
  rules:
    - http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: envoy
                port:
                  number: 30000
`))
		if err != nil {
			return err
		}
		kubeObject.SetAnnotation("config.kubernetes.io/managed-by", fnc.FnUri)
	}
	r.items = append(r.items, kubeObject)
	return nil
}

func (r *resources) ensureConfigMapForEnvoy(rl *fn.ResourceList, routes routes.Routes) error {
	kubeObject, found := common.FindObject(rl, "", "v1", "ConfigMap", "envoy")
	if !found {
		var err error
		kubeObject, err = fn.ParseKubeObject([]byte(`# Generate by gcr.io/treactor/kpt-fn/gateway-api
apiVersion: v1
kind: ConfigMap
metadata:
  name: envoy
  namespace: generated
data:
  envoy.yaml: ""
`))
		if err != nil {
			return err
		}
		kubeObject.SetAnnotation("config.kubernetes.io/managed-by", fnc.FnUri)
	}

	err := kubeObject.SetNestedField(createEnvoyConfig(routes), "data", "envoy.yaml")
	if err != nil {
		return err
	}

	r.items = append(r.items, kubeObject)
	return err
}
