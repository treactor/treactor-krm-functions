package main

import (
	"fmt"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func outputArgoApplication(kptFile *fn.KubeObject, path string) error {
	destinationName := kptFile.GetAnnotation("kpt-repo.argocd.kpt.dev/destination-name")
	projectName := kptFile.GetAnnotation("kpt-repo.argocd.kpt.dev/project-name")
	if destinationName == "" || projectName == "" {
		return nil
	}
	kptFile.GetAnnotation("kpt-repo.argocd.kpt.dev/sync-policy")
	dirPath := filepath.Dir(path)

	name := kptFile.GetName()
	name = namePrefix + strings.ReplaceAll(name, "_", "-")

	ko, err := fn.ParseKubeObject([]byte(`
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: app
  namespace: argocd
spec:
  destination:
    namespace: default
    name: ''
  project: default
  source:
    path: absolute/path
    plugin:
      name: kpt-render
      env:
    repoURL: 'https://github.com/foo/bar.git'
    targetRevision: main
  syncPolicy:
    automated:
      prune: true
      selfHeal: true`))
	if err != nil {
		return err
	}
	ko.SetNestedField(deploymentRepo, "spec", "source", "repoURL")
	ko.SetNestedField(dirPath, "spec", "source", "path")
	ko.SetNestedField(destinationName, "spec", "destination", "name")
	ko.SetNestedField(projectName, "spec", "project")
	ko.SetName(name)
	fmt.Print(ko.String())
	fmt.Println("---")
	return nil
}

var deploymentRepo string
var namePrefix string

func loadConfig() error {
	file, err := os.ReadFile("KptArgoRepo")
	if err != nil {
		return err
	}
	ko, err := fn.ParseKubeObject(file)
	if err != nil {
		return err
	}
	deploymentRepo = ko.GetString("repoUrl")
	namePrefix = ko.GetString("namePrefix")
	return nil
}

func main() {
	err := loadConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.HasSuffix(path, "Kptfile") {
				bytes, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				ko, err := fn.ParseKubeObject(bytes)
				if err != nil {
					return err
				}
				err = outputArgoApplication(ko, path)
				if err != nil {
					return err
				}
			}

			return nil
		})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
