// This file will be processed and embedded to pluginator.

package main

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"os"
)

func Run(rl *fn.ResourceList) (bool, error) {
	return true, nil
}

func main() {
	if err := fn.AsMain(fn.ResourceListProcessorFunc(Run)); err != nil {
		os.Exit(1)
	}
}
