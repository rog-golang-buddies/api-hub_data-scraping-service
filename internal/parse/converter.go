package parse

import (
	"context"
	"errors"

	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/apiSpecDoc"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/fileresource"
)

// Converter converts file data to API specification document using specific file type
//
//go:generate mockgen -source=converter.go -destination=./mocks/converter.go -package=parse
type Converter interface {
	Convert(ctx context.Context, file *fileresource.FileResource) (*apiSpecDoc.ApiSpecDoc, error)
}

type ConverterImpl struct {
	//For instance, we may have a map to hold parsers for different types. And populate it in NewConverter
	parsers map[fileresource.AsdFileType]Parser
}

// Convert gets bytes slice with json/yaml content and a filetype matching the type of the content and returns parsed ApiSpecDoc.
func (c *ConverterImpl) Convert(ctx context.Context, file *fileresource.FileResource) (*apiSpecDoc.ApiSpecDoc, error) {
	//Just example
	parser, ok := c.parsers[file.Type]
	if !ok {
		return nil, errors.New("file type not supported")
	}
	apiSpec, err := parser.Parse(ctx, file.Content)
	if err != nil {
		return nil, err
	}

	return apiSpec, nil
}

func NewConverter(parsers []Parser) Converter {
	parsersMap := make(map[fileresource.AsdFileType]Parser)
	for _, parser := range parsers {
		parsersMap[parser.GetType()] = parser
	}
	return &ConverterImpl{
		parsers: parsersMap,
	}
}
