package queue_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue"
	handler "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/handler/mocks"
	mock_queue "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListenerImpl_Start_noErrors_returnNil(t *testing.T) {
	listener := queue.ListenerImpl{}
	ctrl := gomock.NewController(t)
	consumer := mock_queue.NewMockConsumer(ctrl)
	handl := handler.NewMockHandler(ctrl)
	expectedUrl := "test url"
	consumer.EXPECT().StartConsuming(gomock.Any(), expectedUrl, []string{}, gomock.Any()).Return(nil)
	conf := config.QueueConfig{
		UrlRequestQueue: expectedUrl,
	}
	err := listener.Start(context.Background(), consumer, &conf, handl)
	assert.Nil(t, err)
}

func TestListenerImpl_Start_noStartConsumingError_returnError(t *testing.T) {
	listener := queue.ListenerImpl{}
	ctrl := gomock.NewController(t)
	consumer := mock_queue.NewMockConsumer(ctrl)
	handl := handler.NewMockHandler(ctrl)
	expectedUrl := "test url"
	expectedErr := errors.New("expected unexpected error")
	consumer.EXPECT().StartConsuming(gomock.Any(), expectedUrl, []string{}, gomock.Any()).Return(expectedErr)
	conf := config.QueueConfig{
		UrlRequestQueue: expectedUrl,
	}
	err := listener.Start(context.Background(), consumer, &conf, handl)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err, "method must return error from StartConsuming method")
}

func TestListenerImpl_Start_nilConsumer_returnError(t *testing.T) {
	listener := queue.ListenerImpl{}
	ctrl := gomock.NewController(t)
	handl := handler.NewMockHandler(ctrl)
	expectedUrl := "test url"
	conf := config.QueueConfig{
		UrlRequestQueue: expectedUrl,
	}
	err := listener.Start(context.Background(), nil, &conf, handl)
	assert.NotNil(t, err)
}

func TestListenerImpl_Start_nilConfig_returnError(t *testing.T) {
	listener := queue.ListenerImpl{}
	ctrl := gomock.NewController(t)
	consumer := mock_queue.NewMockConsumer(ctrl)
	handl := handler.NewMockHandler(ctrl)
	err := listener.Start(context.Background(), consumer, nil, handl)
	assert.NotNil(t, err)
}

func TestNewListener(t *testing.T) {
	listener := queue.NewListener()
	assert.NotNil(t, listener)
}
