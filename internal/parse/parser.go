package parse

import (
	"context"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/fileresource"
	"github.com/rog-golang-buddies/api_hub_common/apispecdoc"
)

// Parser is common interface with functionality
// to parse content of the specific API specification document
// and to construct ApiSpecDoc object from it
//
//go:generate mockgen -source=parser.go -destination=./mocks/parser.go -package=parse
type Parser interface {
	//Parse the bytes slice to a ApiSecDoc
	Parse(ctx context.Context, content []byte) (*apispecdoc.ApiSpecDoc, error)

	//GetType returns the type (json or yaml) of the parser
	GetType() fileresource.AsdFileType
}
