package main

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	istiov1beta1 "istio.io/api/networking/v1beta1"
	gwv1beta1 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

type Ingress struct {
	routes []string
}

func (i *Ingress) addRoutes(httpRoute *fn.KubeObject) {

	var rules []gwv1beta1.HTTPRouteRule
	httpRoute.Get(&rules, "spec", "rules")
	for _, rule := range rules {
		for _, match := range rule.Matches {
			i.routes = append(i.routes, *match.Path.Value)
		}
	}

}

func (i *Ingress) getVirtualService() (*fn.KubeObject, error) {
	vs, err := fn.ParseKubeObject([]byte(`# Generate by gcr.io/treactor/kpt-fn/gateway-to-virtualservice
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: vs
  namespace: default
spec:
  gateways:
    - istio-system/istio-gateway-tls-open
  hosts:
    - treactor.example.com
  http: []
`))

	var routes []*istiov1beta1.HTTPRoute

	for _, route := range i.routes {
		routes = append(routes, &istiov1beta1.HTTPRoute{
			Match: []*istiov1beta1.HTTPMatchRequest{
				{
					Uri: &istiov1beta1.StringMatch{
						MatchType: &istiov1beta1.StringMatch_Exact{
							Exact: route,
						},
					},
				},
			},
			Route: []*istiov1beta1.HTTPRouteDestination{{
				Destination: &istiov1beta1.Destination{
					Host: "treactor",
					Port: &istiov1beta1.PortSelector{Number: 3000},
				},
			}},
		})
	}

	vs.SetNestedField(&routes, "spec", "http")

	return vs, err
}
