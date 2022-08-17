package openapi

import (
	"context"
	"github.com/golang/mock/gomock"
	mock_logger "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/logger/mocks"
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
	ctrl := gomock.NewController(t)
	log := mock_logger.NewMockLogger(ctrl)
	content, err := os.ReadFile("./mocks/github_stub.yml")
	if err != nil {
		return
	}
	openAPI, err := parseOpenAPI(ctx, content)
	assert.Nil(t, err)
	if err != nil {
		return
	}
	asd := openapiToApiSpec(log, openAPI)
	assert.NotNil(t, asd)
}
