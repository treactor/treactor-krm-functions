package main

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
)

type CreateAtomsConfig struct {
	MaxNumber     int32           `yaml:"maxNumber,omitempty"`
	MaxBond       int32           `yaml:"maxBond,omitempty"`
	Groups        []GroupConfig   `yaml:"groups,omitempty"`
	FrameworkGo   FrameworkConfig `yaml:"treactorGo,omitempty"`
	FrameworkJava FrameworkConfig `yaml:"treactorJava,omitempty"`
	FrameworkNode FrameworkConfig `yaml:"treactorNode,omitempty"`
}

type GroupConfig struct {
	Group     int    `yaml:"group,omitempty"`
	Framework string `yaml:"framework,omitempty"`
}

type FrameworkConfig struct {
	ImageTag string `yaml:"imageTag,omitempty"`
}

func (ca *CreateAtomsConfig) Config(functionConfig *fn.KubeObject) (error, bool) {
	functionConfig.AsOrDie(ca)
	return nil, false
}
