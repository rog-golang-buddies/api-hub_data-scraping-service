package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseOpenAPI(t *testing.T) {
	ctx := context.Background()
	content, err := os.ReadFile("./mocks/github_stub.yml")
	if err != nil {
		return
	}
	openAPI, err := parseOpenAPI(ctx, content)
	assert.Nil(t, err)
	if err != nil {
		return
	}
	assert.NotNil(t, openAPI)
}

func TestOpenapiToApiSpec(t *testing.T) {
	ctx := context.Background()
	content, err := os.ReadFile("./mocks/github_stub.yml")
	if err != nil {
		return
	}
	openAPI, err := parseOpenAPI(ctx, content)
	assert.Nil(t, err)
	if err != nil {
		return
	}
	asd := openapiToApiSpec(openAPI)
	assert.NotNil(t, asd)
}
