package envoy

import (
	"github.com/treactor/treactor-kpt-functions/gateway-api/pkg/routes"
	"testing"
)

func TestConfig(t *testing.T) {
	r := routes.New()
	backend := r.AddBackend("treactor-api", 80)
	r.AddPrefixRoute(backend, "/treact/reactions")
	r.AddPrefixRoute(backend, "/treact")
	r.AddPrefixRoute(backend, "/treact/info")
	r.AddPrefixRoute(backend, "/treact/nodes/501/info")
	backend = r.AddBackend("bond-n", 80)
	r.AddPrefixRoute(backend, "/treact/nodes/299/info")
	backend = r.AddBackend("atom-h", 80)
	r.AddPrefixRoute(backend, "/treact/nodes/1/info")
	backend = r.AddBackend("treactor-ui", 80)
	r.AddExactRoute(backend, "/")
	r.AddExactRoute(backend, "/index.html")
	r.AddExactRoute(backend, "/robot.txt")
	r.AddPrefixRoute(backend, "/static")
	r.AddPrefixRoute(backend, "/static")
	r.AddPrefixRoute(backend, "/treact/nodes/500/info")
}
