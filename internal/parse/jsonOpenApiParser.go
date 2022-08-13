package parse

import (
	"errors"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/apiSpecDoc"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/fileresource"
)

//JsonOpenApiParser implementation for parsing json open API files
type JsonOpenApiParser struct {
}

func (joap *JsonOpenApiParser) parse(content []byte) (*apiSpecDoc.ApiSpecDoc, error) {
	return nil, errors.New("not implemented")
}

func (joap *JsonOpenApiParser) getType() fileresource.AsdFileType {
	return fileresource.JsonOpenAPI
}

func NewJsonOpenApiParser() Parser {
	return &JsonOpenApiParser{}
}
