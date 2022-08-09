package process

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/fileresource"
	load "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/load/mocks"
	parse "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/parse/mocks"
	recognize "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/recognize/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_RecognizeFail_processReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)

	contentLoader := load.NewMockContentLoader(ctrl)
	recognizer := recognize.NewMockRecognizer(ctrl)
	converter := parse.NewMockConverter(ctrl)

	ctx := context.Background()
	url := "test_url"
	expectedErr := errors.New("recognize error")
	fileResource := new(fileresource.FileResource)

	loadCall := contentLoader.EXPECT().Load(ctx, url).Times(1).Return(fileResource, nil)
	recognizer.EXPECT().RecognizeFileType(fileResource).After(loadCall).Times(1).Return(fileresource.Undefined, expectedErr)

	processor, err := NewProcessor(recognizer, converter, contentLoader)
	assert.Nil(t, err)
	assert.NotNil(t, processor, "Processor must not be nil")

	asd, err := processor.Process(ctx, url)
	assert.Nil(t, asd)
	assert.Equal(t, expectedErr, err, "Should return error from recognizer")
}

func Test_ContentLoaderFail_processReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)

	contentLoader := load.NewMockContentLoader(ctrl)
	recognizer := recognize.NewMockRecognizer(ctrl)
	converter := parse.NewMockConverter(ctrl)

	ctx := context.Background()
	url := "test_url"
	expectedErr := errors.New("contentload error")

	contentLoader.EXPECT().Load(ctx, url).Times(1).Return(nil, expectedErr)

	processor, err := NewProcessor(recognizer, converter, contentLoader)
	assert.Nil(t, err)

	assert.NotNil(t, processor, "Processor must not be nil")

	asd, err := processor.Process(ctx, url)
	assert.Nil(t, asd)
	assert.Equal(t, expectedErr, err, "Should return error from contentLoader")
}

func Test_ConverterFail_processReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)

	contentLoader := load.NewMockContentLoader(ctrl)
	recognizer := recognize.NewMockRecognizer(ctrl)
	converter := parse.NewMockConverter(ctrl)

	ctx := context.Background()
	url := "test_url_from_yaml_openapi_file"
	expectedErr := errors.New("convert error")
	fileResource := new(fileresource.FileResource)

	loadCall := contentLoader.EXPECT().Load(ctx, url).Times(1).Return(fileResource, nil)
	recognizeCall := recognizer.EXPECT().RecognizeFileType(fileResource).After(loadCall).Times(1).Return(fileresource.YamlOpenApi, nil)
	converter.EXPECT().Convert(gomock.Any(), gomock.Any()).Times(1).After(recognizeCall).Return(nil, expectedErr)

	processor, err := NewProcessor(recognizer, converter, contentLoader)
	assert.Nil(t, err)
	assert.NotNil(t, processor, "Processor must not be nil")

	asd, err := processor.Process(ctx, url)
	assert.Nil(t, asd)
	assert.Equal(t, expectedErr, err, "Should return error from converter")
}
