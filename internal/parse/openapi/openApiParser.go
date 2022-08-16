package openapi

import (
	"context"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/apiSpecDoc"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/fileresource"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/logger"
)

type Parser struct {
	log logger.Logger
}

func (p *Parser) Parse(ctx context.Context, content []byte) (*apiSpecDoc.ApiSpecDoc, error) {
	openapi, err := parseOpenAPI(ctx, content)
	if err != nil {
		return nil, err
	}
	return openapiToApiSpec(openapi), nil
}

func (p *Parser) GetType() fileresource.AsdFileType {
	return fileresource.OpenApi
}

func NewOpenApi(log logger.Logger) *Parser {
	return &Parser{log: log}
}
