package openapi

import (
	"errors"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/apiSpecDoc"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/fileresource"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/parse"
)

// YamlOpenApiParser implementation for parsing yml open API files
type YamlOpenApiParser struct {
}

func (yoap *YamlOpenApiParser) Parse(content []byte) (*apiSpecDoc.ApiSpecDoc, error) {
	//loader := openapi3.Loader{IsExternalRefsAllowed: false}
	//doc, err := loader.LoadFromData(content)
	//if err != nil {
	//	return nil, err
	//}
	//doc.Components.
	return nil, errors.New("not implemented")
}

func (yoap *YamlOpenApiParser) GetType() fileresource.AsdFileType {
	return fileresource.JsonOpenAPI
}

func NewYamlOpenApiParser() parse.Parser {
	return &YamlOpenApiParser{}
}
