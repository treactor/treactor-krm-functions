package routes

import (
	"fmt"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	gwv1beta1 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	"sort"
	"strings"
)

type Backend struct {
	Name string
	Host string
	Port int32
}

type Route struct {
	Backend
	Exact bool
	Path  string
}

type Routes struct {
	routes   map[string]Route
	backends map[string]Backend
}

func New() Routes {
	return Routes{
		routes:   map[string]Route{},
		backends: map[string]Backend{},
	}
}

func (r *Routes) AddRoutes(httpRoute *fn.KubeObject) {

	httpRules, _, err := httpRoute.NestedSlice("spec", "rules")
	if err != nil {
		return
	}
	for _, httpRule := range httpRules {
		// The SDK doesn't seem to be able to parse the backendRefs, so working with Object directly
		backendRef := httpRule.GetSlice("backendRefs")
		if len(backendRef) > 0 {
			b := r.AddBackend(backendRef[0].GetString("name"), backendRef[0].GetInt("port"))
			var matches []gwv1beta1.HTTPRouteMatch
			httpRule.Get(&matches, "matches")
			for _, match := range matches {
				pathMatchType := gwv1beta1.PathMatchPathPrefix
				if match.Path.Type != nil {
					pathMatchType = *match.Path.Type
				}
				if pathMatchType == gwv1beta1.PathMatchExact {
					r.AddExactRoute(b, *match.Path.Value)
				} else {
					r.AddPrefixRoute(b, *match.Path.Value)
				}
			}
		}
	}
}

func (r *Routes) AddBackend(host string, port int64) Backend {
	b := Backend{
		Host: host,
		Port: int32(port),
		Name: fmt.Sprintf("%s_%d", strings.Replace(host, "-", "_", -1), port),
	}
	r.backends[b.Name] = b
	return b
}

func (r *Routes) AddExactRoute(backend Backend, path string) {
	route := Route{
		Backend: backend,
		Exact:   true,
		Path:    path,
	}
	r.routes[path] = route
}

func (r *Routes) AddPrefixRoute(backend Backend, path string) {
	route := Route{
		Backend: backend,
		Exact:   false,
		Path:    path,
	}
	r.routes[path] = route
}

func (r *Routes) GetRoutes() []Route {
	v := make([]Route, 0, len(r.routes))
	for _, value := range r.routes {
		v = append(v, value)
	}
	sort.Sort(byRoutePath(v))
	return v
}

func (r *Routes) Backends() []Backend {
	v := make([]Backend, 0, len(r.backends))
	for _, value := range r.backends {
		v = append(v, value)
	}
	sort.Sort(byBackendName(v))
	return v
}

func (r *Routes) RouteCount() int {
	return len(r.routes)
}

func (r *Routes) BackendCount() int {
	return len(r.backends)
}

type byRoutePath []Route

func (s byRoutePath) Len() int {
	return len(s)
}
func (s byRoutePath) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byRoutePath) Less(i, j int) bool {
	if s[i].Exact && !s[j].Exact {
		return true
	} else if !s[i].Exact && s[j].Exact {
		return false
	}
	if len(s[i].Path) == len(s[j].Path) {
		return strings.Compare(s[i].Path, s[j].Path) == -1
	}
	return len(s[i].Path) > len(s[j].Path)
}

type byBackendName []Backend

func (s byBackendName) Len() int {
	return len(s)
}
func (s byBackendName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byBackendName) Less(i, j int) bool {
	return strings.Compare(s[i].Name, s[j].Name) == -1
}
