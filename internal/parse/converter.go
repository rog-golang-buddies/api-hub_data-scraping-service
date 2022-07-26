package parse

import (
	"errors"

	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/apiSpecDoc"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/fileresource"
)

//Converter converts file data to API specification document using specific file type
//go:generate mockgen -source=converter.go -destination=./mocks/converter.go -package=parse
type Converter interface {
	Convert(content []byte, fileType fileresource.AsdFileType) (*apiSpecDoc.ApiSpecDoc, error)
}

type ConverterImpl struct {
	//For instance, we may have a map to hold parsers for different types. And populate it in NewConverter
	parsers map[fileresource.AsdFileType]Parser
}

// Gets bytes slice with json/yaml content and a filetype matching the type of the content and returns parsed ApiSpecDoc.
func (c *ConverterImpl) Convert(content []byte, fileType fileresource.AsdFileType) (*apiSpecDoc.ApiSpecDoc, error) {
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
	parsersMap := make(map[fileresource.AsdFileType]Parser)
	for _, parser := range parsers {
		parsers[parser.getType()] = parser
	}
	return &ConverterImpl{
		parsers: parsersMap,
	}
}
