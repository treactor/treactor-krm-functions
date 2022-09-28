package ingress

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"github.com/treactor/treactor-kpt-functions/common"
	"github.com/treactor/treactor-kpt-functions/gateway-api/pkg/fnc"
	"github.com/treactor/treactor-kpt-functions/gateway-api/pkg/routes"
	networkingv1 "k8s.io/api/networking/v1"
)

type resources struct {
	items []*fn.KubeObject
}

func Create(rl *fn.ResourceList, routes routes.Routes) ([]*fn.KubeObject, error) {
	r := resources{items: []*fn.KubeObject{}}
	err := r.ensureIngress(rl, routes)
	if err != nil {
		return nil, err
	}
	return r.items, nil
}

func (r *resources) ensureIngress(rl *fn.ResourceList, routes routes.Routes) error {
	name := "treactor"
	kubeObject, found := common.FindObject(rl, "networking.k8s.io", "v1", "Ingress", name)
	if !found {
		var err error
		kubeObject, err = fn.ParseKubeObject([]byte(`# Generate by gcr.io/treactor/kpt-fn/gateway-api
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: treactor
`))
		if err != nil {
			return err
		}
	}

	var httpRoutes []networkingv1.HTTPIngressPath
	for _, route := range routes.GetRoutes() {
		pathType := networkingv1.PathTypePrefix
		if route.Exact {
			pathType = networkingv1.PathTypeExact
		}
		httpRoutes = append(httpRoutes, networkingv1.HTTPIngressPath{
			Path:     route.Path,
			PathType: &pathType,
			Backend: networkingv1.IngressBackend{
				Service: &networkingv1.IngressServiceBackend{
					Name: route.Host,
					Port: networkingv1.ServiceBackendPort{
						Number: route.Port,
					},
				},
			},
		})
	}
	err := kubeObject.SetNestedField([]networkingv1.IngressRule{{
		IngressRuleValue: networkingv1.IngressRuleValue{
			HTTP: &networkingv1.HTTPIngressRuleValue{Paths: httpRoutes},
		},
	}}, "spec", "rules")
	if err != nil {
		return err
	}
	kubeObject.SetAnnotation("config.kubernetes.io/managed-by", fnc.FnUri)
	r.items = append(r.items, kubeObject)
	return err
}
