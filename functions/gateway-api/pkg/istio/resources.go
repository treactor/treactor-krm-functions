package istio

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"github.com/treactor/treactor-kpt-functions/common"
	"github.com/treactor/treactor-kpt-functions/gateway-api/pkg/fnc"
	"github.com/treactor/treactor-kpt-functions/gateway-api/pkg/routes"
	istiov1beta1 "istio.io/api/networking/v1beta1"
)

type resources struct {
	items []*fn.KubeObject
}

func Create(rl *fn.ResourceList, routes routes.Routes) ([]*fn.KubeObject, error) {
	r := resources{items: []*fn.KubeObject{}}
	err := r.ensureVirtualService(rl, routes)
	if err != nil {
		return nil, err
	}
	err = r.ensureIstioGateway(rl, routes)
	if err != nil {
		return nil, err
	}
	return r.items, nil
}

func (r *resources) ensureVirtualService(rl *fn.ResourceList, routes routes.Routes) error {
	name := "treactor"
	kubeObject, found := common.FindObject(rl, "networking.istio.io", "v1beta1", "VirtualService", name)
	if !found {
		var err error
		kubeObject, err = fn.ParseKubeObject([]byte(`# Generate by gcr.io/treactor/kpt-fn/gateway-api
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: treactor
  namespace: generated
spec:
  gateways:
    - treactor
  hosts:
    - "*"
  http: []
`))
		if err != nil {
			return err
		}
	}

	var httpRoutes []*istiov1beta1.HTTPRoute
	for _, route := range routes.GetRoutes() {
		var match []*istiov1beta1.HTTPMatchRequest
		if route.Exact {
			match = []*istiov1beta1.HTTPMatchRequest{
				{
					Uri: &istiov1beta1.StringMatch{
						MatchType: &istiov1beta1.StringMatch_Exact{
							Exact: route.Path,
						},
					},
				},
			}
		} else {
			match = []*istiov1beta1.HTTPMatchRequest{
				{
					Uri: &istiov1beta1.StringMatch{
						MatchType: &istiov1beta1.StringMatch_Prefix{
							Prefix: route.Path,
						},
					},
				},
			}
		}
		httpRoutes = append(httpRoutes, &istiov1beta1.HTTPRoute{
			Match: match,
			Route: []*istiov1beta1.HTTPRouteDestination{{
				Destination: &istiov1beta1.Destination{
					Host: route.Host,
					Port: &istiov1beta1.PortSelector{Number: uint32(route.Port)},
				},
			}},
		})
		kubeObject.SetName(name)
		kubeObject.SetAnnotation("config.kubernetes.io/managed-by", fnc.FnUri)
	}

	err := kubeObject.SetNestedField(&httpRoutes, "spec", "http")
	if err != nil {
		return err
	}
	r.items = append(r.items, kubeObject)
	return err
}

func (r *resources) ensureIstioGateway(rl *fn.ResourceList, routes routes.Routes) error {
	name := "treactor"
	kubeObject, found := common.FindObject(rl, "networking.istio.io", "v1beta1", "Gateway", name)
	if !found {
		var err error
		kubeObject, err = fn.ParseKubeObject([]byte(`# Generate by gcr.io/treactor/kpt-fn/gateway-api
apiVersion: networking.istio.io/v1beta1
kind: Gateway
metadata:
  name: treactor
  namespace: generated
spec:
  selector:
    app: istio-ingress
  servers:
`))
		if err != nil {
			return err
		}
		kubeObject.SetName(name)
		kubeObject.SetAnnotation("config.kubernetes.io/managed-by", fnc.FnUri)
	}

	var servers []*istiov1beta1.Server
	servers = append(servers, &istiov1beta1.Server{
		Port: &istiov1beta1.Port{
			Number:   80,
			Protocol: "HTTP",
			Name:     "http",
		},
		Hosts: routes.GetHosts(),
	})

	err := kubeObject.SetNestedField(&servers, "spec", "servers")
	if err != nil {
		return err
	}
	r.items = append(r.items, kubeObject)
	return nil
}
