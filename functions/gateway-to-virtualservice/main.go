// This file will be processed and embedded to pluginator.

package main

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"os"
)

//corev1 "k8s.io/api/core/v1"
//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	fnConfigGroup      = "fn.kpt.dev"
	fnConfigVersion    = "v1alpha1"
	fnConfigAPIVersion = fnConfigGroup + "/" + fnConfigVersion
	fnConfigKind       = "EnsureSA"
)

/*

apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: insights-gateway-chi-local-gateway
  namespace: insights
spec:
  gateways:
    - istio-system/istio-gateway-tls-open
  hosts:
    - insights-app.dev.cp.collibra.dev
  http: []
    - match:
        - uri:
            prefix: /grafana
      retries:
        attempts: 1
        perTryTimeout: 5s
      rewrite:
        uri: ' '
      route:
        - destination:
            host: insights-grafana
            port:
              number: 3000
          weight: 100
      timeout: 10s
    - match:
        - uri:
            prefix: /
      retries:
        attempts: 1
        perTryTimeout: 5s
      route:
        - destination:
            host: dgc-core
            port:
              number: 8080
          weight: 100
      timeout: 10s


*/
func Run(rl *fn.ResourceList) (bool, error) {
	var ingress Ingress
	for _, kubeObject := range rl.Items {
		if kubeObject.IsGVK("gateway.networking.k8s.io/v1alpha2", "HTTPRoute") {
			ingress.addRoutes(kubeObject)
			kubeObject.SetAnnotation("config.kubernetes.io/local-config", "true")
		} else if kubeObject.IsGVK("gateway.networking.k8s.io/v1alpha2", "HTTPRoute") {

		} else {
		}
	}

	vs, err := ingress.getVirtualService()
	if err != nil {
		return false, err
	}
	_ = rl.UpsertObjectToItems(vs, nil, false)

	// This result message will be displayed in the function evaluation time.
	//rl.Results = append(rl.Results, fn.GeneralResult("Ensure ServiceAccount for each `Deployment` resources", fn.Info))
	return true, nil
}

func main() {
	if err := fn.AsMain(fn.ResourceListProcessorFunc(Run)); err != nil {
		os.Exit(1)
	}
}
