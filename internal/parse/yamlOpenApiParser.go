package parse

import (
	"errors"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/apiSpecDoc"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/fileresource"
)

//YamlOpenApiParser implementation for parsing yml open API files
type YamlOpenApiParser struct {
}

func (yoap *YamlOpenApiParser) parse(content []byte) (*apiSpecDoc.ApiSpecDoc, error) {
	return nil, errors.New("not implemented")
}

func (yoap *YamlOpenApiParser) getType() fileresource.AsdFileType {
	return fileresource.JsonOpenAPI
}

func NewYamlOpenApiParser() Parser {
	return &YamlOpenApiParser{}
}
