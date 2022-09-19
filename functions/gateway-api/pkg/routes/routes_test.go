package routes

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	r := New()
	ko, err := fn.ParseKubeObject([]byte(`# Provided
apiVersion: gateway.networking.k8s.io/v1alpha2
kind: HTTPRoute
metadata:
  name: treactor-ui
  annotations:
    config.kubernetes.io/local-config: "true"
spec:
  parentRefs:
    - name: gateway
  hostnames: ["treactor.ioto.pe"]
  rules:
    - backendRefs:
      - name: treactor-ui
        port: 3330
      matches:
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
`))
	assert.NoError(t, err)
	r.AddRoutes(ko)
	assert.Equal(t, 1, len(r.backends))
	assert.Equal(t, 6, len(r.routes))
	sorted := r.GetRoutes()

	assert.Equal(t, "treactor_ui_3330", sorted[0].Name)
	assert.Equal(t, "treactor_ui_3330", sorted[1].Name)
	assert.Equal(t, "treactor_ui_3330", sorted[2].Name)
	assert.Equal(t, "treactor_ui_3330", sorted[3].Name)
	assert.Equal(t, "treactor_ui_3330", sorted[4].Name)
	assert.Equal(t, "treactor_ui_3330", sorted[5].Name)

	assert.Equal(t, true, sorted[0].Exact)
	assert.Equal(t, true, sorted[1].Exact)
	assert.Equal(t, true, sorted[2].Exact)
	assert.Equal(t, false, sorted[3].Exact)
	assert.Equal(t, false, sorted[4].Exact)
	assert.Equal(t, false, sorted[5].Exact)

	assert.Equal(t, "/index.html", sorted[0].Path)
	assert.Equal(t, "/robot.txt", sorted[1].Path)
	assert.Equal(t, "/", sorted[2].Path)
	assert.Equal(t, "/treact/nodes/500/info", sorted[3].Path)
	assert.Equal(t, "/treact/reactions", sorted[4].Path)
	assert.Equal(t, "/static", sorted[5].Path)
}
