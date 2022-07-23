package parse

import (
	"errors"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/model"
)

//Converter converts file data to API specification document using specific file type
//go:generate mockgen -source=converter.go -destination=./mocks/converter.go -package=parse
type Converter interface {
	Convert(content []byte, fileType model.AsdFileType) (*model.ApiSpecDoc, error)
}

type ConverterImpl struct {
	//For instance, we may have a map to hold parsers for different types. And populate it in NewConverter
	parsers map[model.AsdFileType]Parser
}

func (c *ConverterImpl) Convert(content []byte, fileType model.AsdFileType) (*model.ApiSpecDoc, error) {
	//Just example
	parser, ok := c.parsers[fileType]
	if !ok {
		return nil, errors.New("file type not supported")
	}
	apiSpec, err := parser.parse(content)
	if err != nil {
		return nil, err
	}

	return apiSpec, nil
}

func NewConverter(parsers []Parser) Converter {
	parsersMap := make(map[model.AsdFileType]Parser)
	for _, parser := range parsers {
		parsers[parser.getType()] = parser
	}
	return &ConverterImpl{
		parsers: parsersMap,
	}
}
