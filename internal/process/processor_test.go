package process

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/fileresource"
	load "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/load/mocks"
	parse "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/parse/mocks"
	recognize "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/recognize/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RecognizeFail_processReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)

	contentLoader := load.NewMockContentLoader(ctrl)
	recognizer := recognize.NewMockRecognizer(ctrl)
	converter := parse.NewMockConverter(ctrl)

	ctx := context.Background()
	url := "test_url"
	expectedErr := errors.New("load error")
	fileResource := new(fileresource.FileResource)

	loadCall := contentLoader.EXPECT().Load(ctx, url).Times(1).Return(fileResource, nil)
	recognizer.EXPECT().RecognizeFileType(fileResource).After(loadCall).Times(1).Return(fileresource.Undefined, expectedErr)
	converter.EXPECT().Convert(gomock.Any(), gomock.Any()).MaxTimes(0)

	processor, err := NewProcessor(recognizer, converter, contentLoader)
	assert.Nil(t, err)
	assert.NotNil(t, processor, "Processor must not be nil")

	asd, err := processor.process(ctx, url)
	assert.Nil(t, asd)
	assert.Equal(t, expectedErr, err, "Should return error from recognizer")
}
