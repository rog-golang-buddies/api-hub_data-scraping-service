package parse

import (
	"errors"
	"github.com/rog-golang-buddies/internal/model"
)

//JsonOpenApiParser implementation for parsing json open API files
type JsonOpenApiParser struct {
}

func (joap *JsonOpenApiParser) parse(content []byte) (*model.ApiSpecDoc, error) {
	return nil, errors.New("not implemented")
}

func (joap *JsonOpenApiParser) getType() model.AsdFileType {
	return model.JsonOpenAPI
}

func NewJsonOpenApiParser() Parser {
	return &JsonOpenApiParser{}
}
