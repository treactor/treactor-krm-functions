package envoy

import (
	bootstrapv3 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	clusterv3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpointv3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listenerv3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	routev3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	extrouterv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
	hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/treactor/treactor-kpt-functions/gateway-api/pkg/routes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/anypb"
	//type.googleapis.com/envoy.extensions.filters.http.router.v3.Router
)

type envoyConfig struct {
	routes routes.Routes
}

func createEnvoyConfig(r routes.Routes) string {
	envoy := &envoyConfig{routes: r}
	bootstrap := bootstrapv3.Bootstrap{}
	bootstrap.Admin = &bootstrapv3.Admin{
		Address: &v3.Address{
			Address: &v3.Address_SocketAddress{
				SocketAddress: &v3.SocketAddress{
					Address: "0.0.0.0",
					PortSpecifier: &v3.SocketAddress_PortValue{
						PortValue: 9901,
					},
				},
			},
		},
	}

	var cluster []*clusterv3.Cluster
	for _, backend := range r.Backends() {
		cluster = append(cluster, envoy.createCluster(backend))
	}

	bootstrap.StaticResources = &bootstrapv3.Bootstrap_StaticResources{
		Listeners: []*listenerv3.Listener{
			{
				Name: "listener_0",
				Address: &v3.Address{
					Address: &v3.Address_SocketAddress{
						SocketAddress: &v3.SocketAddress{
							Address: "0.0.0.0",
							PortSpecifier: &v3.SocketAddress_PortValue{
								PortValue: 30000,
							},
						},
					},
				},
				FilterChains: []*listenerv3.FilterChain{{

					Filters: []*listenerv3.Filter{{
						Name:       "envoy.filters.network.http_connection_manager",
						ConfigType: &listenerv3.Filter_TypedConfig{TypedConfig: envoy.createHttpConnectionManager()},
					}},
				}},
			},
		},

		Clusters: cluster,
	}
	options := protojson.MarshalOptions{
		UseProtoNames: true,
		Multiline:     true,
	}
	return options.Format(&bootstrap)
}

func (envoy *envoyConfig) createHttpConnectionManager() *any.Any {
	var envoyRoutes []*routev3.Route
	r := envoy.routes.GetRoutes()
	for _, route := range r {
		envoyRoutes = append(envoyRoutes, envoy.createMatch(route))
	}

	routerExtension, err := anypb.New(
		&extrouterv3.Router{})
	if err != nil {
		return nil
	}
	var x = hcm.HttpConnectionManager{
		CodecType:  hcm.HttpConnectionManager_AUTO,
		StatPrefix: "ingress_http",
		RouteSpecifier: &hcm.HttpConnectionManager_RouteConfig{
			RouteConfig: &routev3.RouteConfiguration{
				Name: "local_route",
				VirtualHosts: []*routev3.VirtualHost{{
					Name:    "local_service",
					Domains: []string{"*"},
					Routes:  envoyRoutes,
				}},
			},
		},
		HttpFilters: []*hcm.HttpFilter{
			{
				Name:       "envoy.filters.http.router",
				ConfigType: &hcm.HttpFilter_TypedConfig{TypedConfig: routerExtension},
			},
		},
	}
	a, err := anypb.New(&x)
	if err != nil {
		return nil
	}

	return a
}

func (envoy *envoyConfig) createMatch(route routes.Route) *routev3.Route {
	if route.Exact {
		return envoy.createMatchPath(route.Path, route.Name)
	} else {
		return envoy.createMatchPrefix(route.Path, route.Name)
	}
}

func (envoy *envoyConfig) createMatchPrefix(path, cluster string) *routev3.Route {
	return &routev3.Route{
		Match: &routev3.RouteMatch{
			PathSpecifier: &routev3.RouteMatch_Prefix{
				Prefix: path,
			},
		},
		Action: &routev3.Route_Route{
			Route: &routev3.RouteAction{
				ClusterSpecifier: &routev3.RouteAction_Cluster{
					Cluster: cluster,
				},
			},
		},
	}
}

func (envoy *envoyConfig) createMatchPath(path, cluster string) *routev3.Route {
	return &routev3.Route{
		Match: &routev3.RouteMatch{
			PathSpecifier: &routev3.RouteMatch_Path{
				Path: path,
			},
		},
		Action: &routev3.Route_Route{
			Route: &routev3.RouteAction{
				ClusterSpecifier: &routev3.RouteAction_Cluster{
					Cluster: cluster,
				},
			},
		},
	}
}

func (envoy *envoyConfig) createCluster(backend routes.Backend) *clusterv3.Cluster {
	//name := envoy.clusterFromHost(host)

	return &clusterv3.Cluster{
		Name: backend.Name,
		ClusterDiscoveryType: &clusterv3.Cluster_Type{
			Type: clusterv3.Cluster_STRICT_DNS,
		},
		ConnectTimeout: &duration.Duration{
			Seconds: 0,
			Nanos:   250000000,
		},
		LbPolicy: clusterv3.Cluster_ROUND_ROBIN,
		LoadAssignment: &endpointv3.ClusterLoadAssignment{
			ClusterName: backend.Name,
			Endpoints: []*endpointv3.LocalityLbEndpoints{
				{
					LbEndpoints: []*endpointv3.LbEndpoint{
						{
							HostIdentifier: &endpointv3.LbEndpoint_Endpoint{
								Endpoint: &endpointv3.Endpoint{
									Address: &v3.Address{
										Address: &v3.Address_SocketAddress{
											SocketAddress: &v3.SocketAddress{
												Address:       backend.Host,
												PortSpecifier: &v3.SocketAddress_PortValue{PortValue: uint32(backend.Port)},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
