package recognize

import (
	"errors"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/logger"

	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/fileresource"
)

// Recognizer provide functionality to recognize file type by content
//
//go:generate mockgen -source=recognizer.go -destination=./mocks/recognizer.go -package=recognize
type Recognizer interface {
	//RecognizeFileType recognizes type of the file by content. Probably we may combine it with validation
	//Also not sure name is needed here. Better to recognize by content (check is file yaml, if yaml - check version;
	//if json - check openApi version) But it is easier to use file extension as starting point to check content.
	RecognizeFileType(resource *fileresource.FileResource) (fileresource.AsdFileType, error)
}

type RecognizerImpl struct {
	log logger.Logger
}

func (r *RecognizerImpl) RecognizeFileType(resource *fileresource.FileResource) (fileresource.AsdFileType, error) {
	r.log.Info("start file '%s' recognizing", resource.Link)

	return fileresource.Undefined, errors.New("not implemented")
}

func NewRecognizer(log logger.Logger) Recognizer {
	return &RecognizerImpl{
		log: log,
	}
}
