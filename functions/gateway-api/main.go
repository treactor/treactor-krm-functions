// This file will be processed and embedded to pluginator.

package main

import (
	"fmt"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"github.com/treactor/treactor-kpt-functions/common/constants"
	"github.com/treactor/treactor-kpt-functions/gateway-api/pkg/config"
	"github.com/treactor/treactor-kpt-functions/gateway-api/pkg/envoy"
	"github.com/treactor/treactor-kpt-functions/gateway-api/pkg/fnc"
	"github.com/treactor/treactor-kpt-functions/gateway-api/pkg/ingress"
	"github.com/treactor/treactor-kpt-functions/gateway-api/pkg/routes"
	"github.com/treactor/treactor-kpt-functions/gateway-api/pkg/virtualservice"
	"os"
)

//corev1 "k8s.io/api/core/v1"
//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func Run(rl *fn.ResourceList) (bool, error) {
	var cfg config.GatewayApiConfig
	err, _ := cfg.Config(rl.FunctionConfig)
	if err != nil {
		os.Exit(1)
	}

	var r = routes.New()

	for _, kubeObject := range rl.Items {
		if kubeObject.IsGVK("gateway.networking.k8s.io", "v1alpha2", "HTTPRoute") {
			r.AddRoutes(kubeObject)
			kubeObject.SetAnnotation(constants.K8sAnnLocalConfig, "true")
		} else if kubeObject.IsGVK("gateway.networking.k8s.io", "v1alpha2", "Gateway") {
			kubeObject.SetAnnotation(constants.K8sAnnLocalConfig, "true")
		}
	}

	var items []*fn.KubeObject
	if cfg.Output == "istio" {
		fnResources, err := virtualservice.Create(rl, r)
		if err != nil {
			return false, err
		}
		items = append(items, fnResources...)
	} else if cfg.Output == "envoy" {
		fnResources, err := envoy.Create(rl, cfg.Envoy, r)
		if err != nil {
			return false, err
		}
		items = append(items, fnResources...)
	} else if cfg.Output == "ingress" {
		fnResources, err := ingress.Create(rl, r)
		if err != nil {
			return false, err
		}
		items = append(items, fnResources...)
	}

	// Collect the unmanaged resources
	for _, item := range rl.Items {
		if item.GetAnnotation("config.kubernetes.io/managed-by") != fnc.FnUri {
			items = append(items, item)
		}
	}
	rl.Items = items

	// This result message will be displayed in the function evaluation time.
	rl.Results = append(rl.Results, fn.GeneralResult(fmt.Sprintf("Added %d backends", r.BackendCount()), fn.Info))
	rl.Results = append(rl.Results, fn.GeneralResult(fmt.Sprintf("Added %d routes", r.RouteCount()), fn.Info))
	return true, nil
}

func main() {
	if err := fn.AsMain(fn.ResourceListProcessorFunc(Run)); err != nil {
		os.Exit(1)
	}
}
