package config

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
)

type GatewayApiConfig struct {
	Output  string         `yaml:"output"`
	Istio   *IstioConfig   `yaml:"istio,omitempty"`
	Ingress *IngressConfig `yaml:"ingress,omitempty"`
	Envoy   *EnvoyConfig   `yaml:"envoy,omitempty"`
}

type IstioConfig struct {
}

type IngressConfig struct {
}

type EnvoyConfig struct {
	EnableIngress bool `yaml:"enableIngress,omitempty"`
}

func (config *GatewayApiConfig) Config(functionConfig *fn.KubeObject) (error, bool) {
	functionConfig.AsOrDie(config)
	return nil, false
}
