package main

import (
	"fmt"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"log"
	"os"
	"path/filepath"
)

func main() {
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				bytes, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				ko, err := fn.ParseKubeObject(bytes)
				if err != nil {
					// ignore non KRM files
					return nil
				}
				if "true" != ko.GetAnnotation("config.kubernetes.io/local-config") {
					fmt.Print(ko.String())
					fmt.Println("---")
				}
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
