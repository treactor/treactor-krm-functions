package main

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnvoy(t *testing.T) {
	resourceTestFile, err := os.ReadFile("data/main_test_envoy.yaml")
	assert.NoError(t, err)
	resourceList, err := fn.ParseResourceList(resourceTestFile)
	run, err := Run(resourceList)
	assert.NoError(t, err)
	assert.True(t, run)
	assert.NoError(t, fn.CheckResourceDuplication(resourceList))
	assert.Equal(t, 4, len(resourceList.Items))
	assert.Equal(t, resourceList.Results[0].Message, "Added 1 backends")
	assert.Equal(t, resourceList.Results[1].Message, "Added 6 routes")
}

func TestEnvoyUpdateDeployment(t *testing.T) {
	resourceTestFile, err := os.ReadFile("data/main_test_envoy_update_deployment.yaml")
	assert.NoError(t, err)
	resourceList, err := fn.ParseResourceList(resourceTestFile)
	run, err := Run(resourceList)
	assert.NoError(t, err)
	assert.True(t, run)
	assert.NoError(t, fn.CheckResourceDuplication(resourceList))
	assert.Equal(t, 4, len(resourceList.Items))
	assert.Equal(t, resourceList.Results[0].Message, "Added 1 backends")
	assert.Equal(t, resourceList.Results[1].Message, "Added 6 routes")
}

func TestEnvoyWithIngress(t *testing.T) {
	resourceTestFile, err := os.ReadFile("data/main_test_envoy_with_ingress.yaml")
	assert.NoError(t, err)
	resourceList, err := fn.ParseResourceList(resourceTestFile)
	run, err := Run(resourceList)
	assert.NoError(t, err)
	assert.True(t, run)
	assert.NoError(t, fn.CheckResourceDuplication(resourceList))
	assert.Equal(t, 5, len(resourceList.Items))
	assert.Equal(t, resourceList.Results[0].Message, "Added 1 backends")
	assert.Equal(t, resourceList.Results[1].Message, "Added 6 routes")
}

func TestIstio(t *testing.T) {
	resourceTestFile, err := os.ReadFile("data/main_test_istio.yaml")
	assert.NoError(t, err)
	resourceList, err := fn.ParseResourceList(resourceTestFile)
	run, err := Run(resourceList)
	assert.NoError(t, err)
	assert.True(t, run)
	assert.NoError(t, fn.CheckResourceDuplication(resourceList))
	assert.Equal(t, 2, len(resourceList.Items))
	assert.Equal(t, resourceList.Results[0].Message, "Added 1 backends")
	assert.Equal(t, resourceList.Results[1].Message, "Added 6 routes")
}
