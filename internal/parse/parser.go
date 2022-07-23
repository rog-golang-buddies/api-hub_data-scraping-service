package parse

import (
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/apiSpecDoc"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/fileresource"
)

//Parser is common interface with functionality
//to parse content of the specific API specification document
//and to construct ApiSpecDoc object from it
type Parser interface {
	parse(content []byte) (*apiSpecDoc.ApiSpecDoc, error)

	getType() fileresource.AsdFileType
}
