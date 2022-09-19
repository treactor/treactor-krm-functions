package common

import "github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"

func FindObject(rl *fn.ResourceList, group, version, kind, name string) (*fn.KubeObject, bool) {
	for _, kubeObject := range rl.Items {

		if kubeObject.IsGVK(group, version, kind) && kubeObject.GetName() == name {
			return kubeObject, true
		}
	}
	return nil, false
}
