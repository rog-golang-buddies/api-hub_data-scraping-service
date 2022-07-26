package process

import (
	"context"

	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/apiSpecDoc"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/load"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/parse"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/recognize"
)

//UrlProcessor represents provide entrypoint for the url processing
//full processing of the incoming
//go:generate mockgen -source=processor.go -destination=./mocks/processor.go -package=process
type UrlProcessor interface {
	process(ctx context.Context, url string) (*apiSpecDoc.ApiSpecDoc, error)
}

type ProcessorImpl struct {
	recognizer    recognize.Recognizer
	converter     parse.Converter
	contentLoader load.ContentLoader
}

// Gets the url of a OpenApi file (Swagger file) string as parameter and returns an
func (p *ProcessorImpl) process(ctx context.Context, url string) (*apiSpecDoc.ApiSpecDoc, error) {
	//Check availability of url
	//...

	//Load content by url. Ctx check is done inside Load function if it's cancelled, returns an error.
	file, err := p.contentLoader.Load(ctx, url)
	if err != nil {
		return nil, err
	}

	//If no errs recognize file type by content
	fileType, err := p.recognizer.RecognizeFileType(file)
	if err != nil {
		return nil, err
	}

	//Parse API spec of defined type
	apiSpec, err := p.converter.Convert(file.Content, fileType)
	if err != nil {
		return nil, err
	}

	return apiSpec, nil
}

func NewProcessor(r recognize.Recognizer, c parse.Converter, cl load.ContentLoader) (UrlProcessor, error) {
	return &ProcessorImpl{
		recognizer:    r,
		converter:     c,
		contentLoader: cl,
	}, nil
}
