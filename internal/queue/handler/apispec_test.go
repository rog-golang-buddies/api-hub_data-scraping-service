package handler

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/dto/apiSpecDoc"
	process "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/process/mocks"
	publisher "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/publisher/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/wagslane/go-rabbitmq"
	"testing"
)

func TestApiSpecDocHandler_Handle_wrongBody_NackDiscard(t *testing.T) {
	ctrl := gomock.NewController(t)
	pub := publisher.NewMockPublisher(ctrl)
	proc := process.NewMockUrlProcessor(ctrl)
	conf := config.QueueConfig{}

	handl := NewApiSpecDocHandler(pub, conf, proc)
	wrongBody := "wrong body"
	delivery := rabbitmq.Delivery{
		Delivery: amqp091.Delivery{Body: []byte(wrongBody)},
	}
	action := handl.Handle(context.Background(), delivery)
	assert.Equal(t, rabbitmq.NackDiscard, action)
}

func TestApiSpecDocHandler_Handle_publishError_NackDiscard(t *testing.T) {
	ctrl := gomock.NewController(t)
	pub := publisher.NewMockPublisher(ctrl)
	proc := process.NewMockUrlProcessor(ctrl)

	queueName := "test queue"
	conf := config.QueueConfig{
		ScrapingResultQueue: queueName,
	}
	proc.EXPECT().Process(gomock.Any(), "test url").Times(1).Return(&apiSpecDoc.ApiSpecDoc{}, nil)
	pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{queueName}), gomock.Any()).Times(1).
		Return(errors.New("publish error"))

	handl := NewApiSpecDocHandler(pub, conf, proc)
	body := `{"FileUrl":"test url","IsNotifyUser":false}`
	delivery := rabbitmq.Delivery{
		Delivery: amqp091.Delivery{Body: []byte(body)},
	}
	action := handl.Handle(context.Background(), delivery)
	assert.Equal(t, rabbitmq.NackDiscard, action)
}

func TestApiSpecDocHandler_Handle_allCorrectNotificationFalse_called1TimeAck(t *testing.T) {
	ctrl := gomock.NewController(t)
	pub := publisher.NewMockPublisher(ctrl)
	proc := process.NewMockUrlProcessor(ctrl)

	queueName := "test queue"
	conf := config.QueueConfig{
		ScrapingResultQueue: queueName,
	}
	proc.EXPECT().Process(gomock.Any(), "test url").Times(1).Return(&apiSpecDoc.ApiSpecDoc{}, nil)
	pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{queueName}), gomock.Any()).Times(1).Return(nil)

	handl := NewApiSpecDocHandler(pub, conf, proc)
	body := `{"FileUrl":"test url","IsNotifyUser":false}`
	delivery := rabbitmq.Delivery{
		Delivery: amqp091.Delivery{Body: []byte(body)},
	}
	action := handl.Handle(context.Background(), delivery)
	assert.Equal(t, rabbitmq.Ack, action)
}

func TestApiSpecDocHandler_Handle_allCorrectNotificationFalse_called2TimesAck(t *testing.T) {
	ctrl := gomock.NewController(t)
	pub := publisher.NewMockPublisher(ctrl)
	proc := process.NewMockUrlProcessor(ctrl)

	resQName := "test queue"
	notQName := "test notification queue"
	conf := config.QueueConfig{
		ScrapingResultQueue: resQName,
		NotificationQueue:   notQName,
	}
	proc.EXPECT().Process(gomock.Any(), "test url").Times(1).Return(&apiSpecDoc.ApiSpecDoc{}, nil)
	firstCall := pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{resQName}), gomock.Any()).Times(1).Return(nil)
	pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{notQName}), gomock.Any()).Times(1).Return(nil).After(firstCall)

	handl := NewApiSpecDocHandler(pub, conf, proc)
	body := `{"FileUrl":"test url","IsNotifyUser":true}`
	delivery := rabbitmq.Delivery{
		Delivery: amqp091.Delivery{Body: []byte(body)},
	}
	action := handl.Handle(context.Background(), delivery)
	assert.Equal(t, rabbitmq.Ack, action)
}

func TestApiSpecDocHandler_Handle_notificationError_called2TimesAck(t *testing.T) {
	ctrl := gomock.NewController(t)
	pub := publisher.NewMockPublisher(ctrl)
	proc := process.NewMockUrlProcessor(ctrl)

	resQName := "test queue"
	notQName := "test notification queue"
	conf := config.QueueConfig{
		ScrapingResultQueue: resQName,
		NotificationQueue:   notQName,
	}
	proc.EXPECT().Process(gomock.Any(), "test url").Times(1).Return(&apiSpecDoc.ApiSpecDoc{}, nil)
	firstCall := pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{resQName}), gomock.Any()).
		Times(1).
		Return(nil)
	pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{notQName}), gomock.Any()).Times(1).
		Return(errors.New("unexpected notification error")).
		After(firstCall)

	handl := NewApiSpecDocHandler(pub, conf, proc)
	body := `{"FileUrl":"test url","IsNotifyUser":true}`
	delivery := rabbitmq.Delivery{
		Delivery: amqp091.Delivery{Body: []byte(body)},
	}
	action := handl.Handle(context.Background(), delivery)
	assert.Equal(t, rabbitmq.Ack, action)
}
