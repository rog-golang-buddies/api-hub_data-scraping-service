package recognize

import (
	"errors"
	"github.com/rog-golang-buddies/internal/model"
)

//Recognizer provide functionality to recognize file type by content
//go:generate mockgen -source=recognizer.go -destination=./mocks/recognizer.go -package=recognize
type Recognizer interface {
	//RecognizeFileType recognizes type of the file by content. Probably we may combine it with validation
	//Also not sure name is needed here. Better to recognize by content (check is file yaml, if yaml - check version;
	//if json - check openApi version) But it is easier to use file extension as starting point to check content.
	RecognizeFileType(resource *model.FileResource) (model.AsdFileType, error)
}

type RecognizerImpl struct {
}

func (r *RecognizerImpl) RecognizeFileType(resource *model.FileResource) (model.AsdFileType, error) {
	return model.Undefined, errors.New("not implemented")
}

func NewRecognizer() Recognizer {
	return &RecognizerImpl{}
}
