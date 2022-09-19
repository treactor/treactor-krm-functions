package virtualservice

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
	return r.items, nil
}

func (r *resources) ensureVirtualService(rl *fn.ResourceList, routes routes.Routes) error {
	name := "treactor"
	kubeObject, found := common.FindObject(rl, "networking.istio.io", "v1alpha3", "VirtualService", name)
	if !found {
		var err error
		kubeObject, err = fn.ParseKubeObject([]byte(`# Generate by gcr.io/treactor/kpt-fn/gateway-api
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: kubeObject
spec:
  gateways:
    - istio-system/ingressgateway
  hosts:
    - treactor.example.com
  http: []
`))
		if err != nil {
			return err
		}
	}

	var httpRoutes []*istiov1beta1.HTTPRoute
	for _, route := range routes.GetRoutes() {
		httpRoutes = append(httpRoutes, &istiov1beta1.HTTPRoute{
			Match: []*istiov1beta1.HTTPMatchRequest{
				{
					Uri: &istiov1beta1.StringMatch{
						MatchType: &istiov1beta1.StringMatch_Exact{
							Exact: route.Path,
						},
					},
				},
			},
			Route: []*istiov1beta1.HTTPRouteDestination{{
				Destination: &istiov1beta1.Destination{
					Host: route.Host,
					Port: &istiov1beta1.PortSelector{Number: uint32(route.Port)},
				},
			}},
		})
		kubeObject.SetAnnotation("config.kubernetes.io/managed-by", fnc.FnUri)
	}

	err := kubeObject.SetNestedField(&httpRoutes, "spec", "http")
	if err != nil {
		return err
	}
	r.items = append(r.items, kubeObject)
	return err
}
