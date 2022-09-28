// This file will be processed and embedded to pluginator.

package main

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"github.com/treactor/treactor-kpt-functions/common"
	"github.com/treactor/treactor-kpt-functions/common/constants"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
)

const (
	fnUri = constants.FnUriPrefix + "ensure-sa"
)

type EnsureSaFn struct {
	items []*fn.KubeObject
}

func CreateFn() *EnsureSaFn {
	function := EnsureSaFn{
		items: []*fn.KubeObject{},
	}
	return &function
}

func (function *EnsureSaFn) ensureServiceAccount(rl *fn.ResourceList, name, namespace string) error {
	kubeObject, found := common.FindObject(rl, "", "v1", "ServiceAccount", name)
	if !found {
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
		var err error
		kubeObject, err = fn.NewFromTypedObject(sa)
		if err != nil {
			return err
		}
		kubeObject.SetAnnotation("config.kubernetes.io/managed-by", fnUri)
	}
	function.items = append(function.items, kubeObject)
	return nil
}

// Collect unmanaged resources
func (function *EnsureSaFn) collectUnmanaged(rl *fn.ResourceList) {
	for _, item := range rl.Items {
		if item.GetAnnotation("config.kubernetes.io/managed-by") != fnUri {
			function.items = append(function.items, item)
		}
	}
}

func Run(rl *fn.ResourceList) (bool, error) {
	function := CreateFn()
	for _, kubeObject := range rl.Items {
		if kubeObject.IsGVK("apps", "v1", "Deployment") {
			name := kubeObject.GetName()
			namespace := kubeObject.GetNamespace()
			kubeObject.SetNestedField(name, "spec", "template", "spec", "serviceAccountName")
			err := function.ensureServiceAccount(rl, name, namespace)
			if err != nil {
				return false, err
			}
		}
	}
	function.collectUnmanaged(rl)
	rl.Items = function.items

	// This result message will be displayed in the function evaluation time.
	rl.Results = append(rl.Results, fn.GeneralResult("Ensure ServiceAccount for each `Deployment` resources", fn.Info))
	return true, nil
}

func main() {
	if err := fn.AsMain(fn.ResourceListProcessorFunc(Run)); err != nil {
		os.Exit(1)
	}
}
