package main

import (
	"fmt"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"github.com/treactor/treactor-kpt-functions/create-atoms/element"
)

type CreateAtomsConfig struct {
	MaxNumber  int32             `yaml:"maxNumber,omitempty"`
	MaxBond    int32             `yaml:"maxBond,omitempty"`
	Groups     []GroupConfig     `yaml:"groups,omitempty"`
	Numbers    []NumberConfig    `yaml:"numbers,omitempty"`
	Frameworks []FrameworkConfig `yaml:"frameworks,omitempty"`

	frameworks map[string]FrameworkConfig
	groups     map[int]FrameworkConfig
	numbers    map[int]FrameworkConfig
}

type GroupConfig struct {
	Group     int    `yaml:"group,omitempty"`
	Framework string `yaml:"framework,omitempty"`
}

type NumberConfig struct {
	Number    int    `yaml:"number,omitempty"`
	Framework string `yaml:"framework,omitempty"`
}

type FrameworkConfig struct {
	Name     string `yaml:"name,omitempty"`
	Image    string `yaml:"image,omitempty"`
	ImageTag string `yaml:"imageTag,omitempty"`
}

func (config *CreateAtomsConfig) Config(functionConfig *fn.KubeObject) (error, bool) {
	functionConfig.AsOrDie(config)
	config.frameworks = map[string]FrameworkConfig{}
	config.groups = map[int]FrameworkConfig{}
	config.numbers = map[int]FrameworkConfig{}
	for _, framework := range config.Frameworks {
		config.frameworks[framework.Name] = framework
	}
	for _, group := range config.Groups {
		config.groups[group.Group] = config.frameworks[group.Framework]
	}
	for _, number := range config.Numbers {
		config.numbers[number.Number] = config.frameworks[number.Framework]
	}
	return nil, false
}

func (config *CreateAtomsConfig) framework(atom element.Element) FrameworkConfig {
	framework, found := config.numbers[int(atom.Number)]
	if found {
		return framework
	}
	framework, found = config.groups[int(atom.Group)]
	if found {
		return framework
	}
	if atom.Number == 500 {
		return config.frameworks["treactorNode"]
	}
	return config.frameworks["treactorGo"]
}

func (config *CreateAtomsConfig) ImageForAtom(atom element.Element) string {
	framework := config.framework(atom)
	return fmt.Sprintf("%s:%s", framework.Image, framework.ImageTag)
}

func (config *CreateAtomsConfig) ImageTagForAtom(atom element.Element) string {
	framework := config.framework(atom)
	return framework.ImageTag
}
