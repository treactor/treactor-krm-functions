package main

import (
	"fmt"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"github.com/treactor/treactor-kpt-functions/create-atoms/element"
	"testing"
)

func TestAddDeploymentForElement(t *testing.T) {
	resourceList := fn.ResourceList{}
	function := CreateFn()
	function.addDeploymentForElement(&resourceList, element.Element{
		Number: 42,
		Symbol: "Mo",
		Name:   "Molybdenum",
	}, CreateAtomsConfig{})
	fmt.Println(function.items)
}

func TestAddRouteForElement(t *testing.T) {
	resourceList := fn.ResourceList{}
	function := CreateFn()
	function.addRouteForElement(&resourceList, element.Element{
		Number: 42,
		Symbol: "Mo",
		Name:   "Molybdenum",
	}, CreateAtomsConfig{})
	fmt.Println(function.items)
}
