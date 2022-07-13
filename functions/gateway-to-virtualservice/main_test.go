package main

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"testing"
)

func TestAddDeploymentForElement(t *testing.T) {

	//fn.NewFromTypedObject()
	kubeObject, _ := fn.ParseKubeObject([]byte(`apiVersion: apps/v1
apiVersion: gateway.networking.k8s.io/v1alpha2
kind: HTTPRoute
metadata:
  name: treactor-ui
  namespace: default
spec:
  parentRefs:
    - name: gateway
      namespace: treactor
  hostnames: ["treactor.ioto.pe"]
  rules:
    - matches:
        - path:
            type: Exact
            value: /
        - path:
            type: Exact
            value: /index.html
        - path:
            type: Exact
            value: /robot.txt
        - path:
            type: PathPrefix
            value: /static
        - path:
            type: PathPrefix
            value: /treact/reactions
        - path:
            type: PathPrefix
            value: /treact/nodes/500/info
      backendRefs:
        - name: treactor-ui
          port: 3330
`))

	var ingress Ingress

	ingress.addRoutes(kubeObject)
	ingress.getVirtualService()
}
