// This file will be processed and embedded to pluginator.

package main

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
)

const (
	fnConfigGroup      = "fn.kpt.dev"
	fnConfigVersion    = "v1alpha1"
	fnConfigAPIVersion = fnConfigGroup + "/" + fnConfigVersion
	fnConfigKind       = "EnsureSA"
)

func addServiceAccountForElement(rl *fn.ResourceList, name, namespace string) error {
	sa := &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
	return rl.UpsertObjectToItems(sa, nil, false)
}

func Run(rl *fn.ResourceList) (bool, error) {
	for _, kubeObject := range rl.Items {
		if kubeObject.IsGVK("apps/v1", "Deployment") {
			name := kubeObject.GetName()
			namespace := kubeObject.GetNamespace()
			kubeObject.SetNestedField(name, "spec", "serviceAccountName")
			err := addServiceAccountForElement(rl, name, namespace)
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
