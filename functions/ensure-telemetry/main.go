// This file will be processed and embedded to pluginator.

package main

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	//corev1 "k8s.io/api/core/v1"
	"os"
)

func upsert(env []map[string]interface{}, envVar map[string]interface{}) []map[string]interface{} {
	name, found := envVar["name"]
	if !found {
		return env
	}
	for ix, e := range env {
		if name == e["name"] {
			env[ix] = envVar
			return env
		}
	}
	return append(env, envVar)
}

func EnsureTelemetryEnv(container *fn.SubObject) error {

	var env []map[string]interface{}
	_, err2 := container.Get(&env, "env")
	if err2 != nil {
		return err2
	}

	env = upsert(env, map[string]interface{}{
		"name": "K8S_POD_NAME",
		"valueFrom": map[string]interface{}{
			"fieldRef": map[string]interface{}{
				"fieldPath": "metadata.name",
			},
		},
	})
	env = upsert(env, map[string]interface{}{
		"name": "K8S_NAMESPACE_NAME",
		"valueFrom": map[string]interface{}{
			"fieldRef": map[string]interface{}{
				"fieldPath": "metadata.namespace",
			},
		},
	})
	env = upsert(env, map[string]interface{}{
		"name": "K8S_POD_UID",
		"valueFrom": map[string]interface{}{
			"fieldRef": map[string]interface{}{
				"fieldPath": "metadata.uid",
			},
		},
	})
	env = upsert(env, map[string]interface{}{
		"name": "HOST_IP",
		"valueFrom": map[string]interface{}{
			"fieldRef": map[string]interface{}{
				"fieldPath": "status.hostIP",
			},
		},
	})
	env = upsert(env, map[string]interface{}{
		"name":  "OTEL_EXPORTER",
		"value": "otlp",
	})
	env = upsert(env, map[string]interface{}{
		"name":  "OTEL_EXPORTER_OTLP_ENDPOINT",
		"value": "http://$(HOST_IP):4317",
	})
	env = upsert(env, map[string]interface{}{
		"name":  "OTEL_EXPORTER_OTLP_PROTOCOL",
		"value": "grpc",
	})
	container.SetNestedField(env, "env")

	return nil
}

func Run(rl *fn.ResourceList) (bool, error) {
	for _, kubeObject := range rl.Items {
		if kubeObject.IsGVK("apps", "v1", "Deployment") {
			//name := kubeObject.GetName()

			slice, _, err := kubeObject.NestedSlice("spec", "template", "spec", "containers")
			if err != nil {
				return false, err
			}

			err = EnsureTelemetryEnv(slice[0])
			if err != nil {
				return false, err
			}
		}
	}

	// This result message will be displayed in the function evaluation time.
	rl.Results = append(rl.Results, fn.GeneralResult("Ensure ServiceAccount for each `Deployment` resources", fn.Info))
	return true, nil
}

func main() {
	if err := fn.AsMain(fn.ResourceListProcessorFunc(Run)); err != nil {
		os.Exit(1)
	}
}
