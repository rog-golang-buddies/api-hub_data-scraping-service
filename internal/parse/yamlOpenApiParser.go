package parse

import (
	"errors"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/model"
)

//YamlOpenApiParser implementation for parsing yml open API files
type YamlOpenApiParser struct {
}

func (yoap *YamlOpenApiParser) parse(content []byte) (*model.ApiSpecDoc, error) {
	return nil, errors.New("not implemented")
}

func (yoap *YamlOpenApiParser) getType() model.AsdFileType {
	return model.JsonOpenAPI
}

func NewYamlOpenApiParser() Parser {
	return &YamlOpenApiParser{}
}
