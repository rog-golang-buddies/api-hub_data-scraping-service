package parse

import "github.com/rog-golang-buddies/internal/model"

//Parser is common interface with functionality
//to parse content of the specific API specification document
//and to construct ApiSpecDoc object from it
type Parser interface {
	parse(content []byte) (*model.ApiSpecDoc, error)

	getType() model.AsdFileType
}
