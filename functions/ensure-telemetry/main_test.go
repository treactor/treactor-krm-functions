package main

import (
	"fmt"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"testing"
)

func TestAddDeploymentForElement(t *testing.T) {

	//fn.NewFromTypedObject()
	kubeObject, _ := fn.ParseKubeObject([]byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: myDeployment
spec:
  replicas: 3
  template:
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
`))

	slice, _, _ := kubeObject.NestedSlice("spec", "template", "spec", "containers")

	EnsureTelemetryEnv(slice[0])
	fmt.Println(kubeObject)
}
