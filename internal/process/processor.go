package process

import (
	"context"
	"errors"

	"github.com/rog-golang-buddies/internal/load"
	"github.com/rog-golang-buddies/internal/model"
	"github.com/rog-golang-buddies/internal/parse"
	"github.com/rog-golang-buddies/internal/recognize"
)

//UrlProcessor represents provide entrypoint for the url processing
//full processing of the incoming
type UrlProcessor interface {
	process(ctx context.Context, url string) (*model.ApiSpecDoc, error)
}

type ProcessorImpl struct {
	recognizer    recognize.Recognizer
	converter     parse.Converter
	contentLoader load.ContentLoader
}

func (p *ProcessorImpl) process(ctx context.Context, url string) (*model.ApiSpecDoc, error) {
	//Check availability of url
	//...

	content := make(chan struct {
		file *model.FileResource
		err  error
	}, 1)

	//Load content by url
	go func() {
		file, err := p.contentLoader.Load(ctx, url)
		content <- struct {
			file *model.FileResource
			err  error
		}{file, err}
	}()

	select {
	case result := <-content:

		if result.err != nil {
			return nil, result.err
		}

		//If no errs recognize file type by content
		fileType, err := p.recognizer.RecognizeFileType(result.file)
		if err != nil {
			return nil, err
		}

		//Parse API spec of defined type
		apiSpec, err := p.converter.Convert(result.file.Content, fileType)
		if err != nil {
			return nil, err
		}

		return apiSpec, nil

	case <-ctx.Done():
		return nil, errors.New("Load cancelled")
	}
}

func NewProcessor(r recognize.Recognizer, c parse.Converter, cl load.ContentLoader) (UrlProcessor, error) {
	return &ProcessorImpl{
		recognizer:    r,
		converter:     c,
		contentLoader: cl,
	}, nil
}
