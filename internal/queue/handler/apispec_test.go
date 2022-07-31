package handler

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	publisher "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/publisher/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/wagslane/go-rabbitmq"
	"testing"
)

func TestApiSpecDocHandler_Handle_wrongBody_NackDiscard(t *testing.T) {
	ctrl := gomock.NewController(t)
	pub := publisher.NewMockPublisher(ctrl)
	conf := config.QueueConfig{}

	handl := NewApiSpecDocHandler(pub, conf)
	wrongBody := "wrong body"
	delivery := rabbitmq.Delivery{
		Delivery: amqp091.Delivery{Body: []byte(wrongBody)},
	}
	action := handl.Handle(delivery)
	assert.Equal(t, rabbitmq.NackDiscard, action)
}

func TestApiSpecDocHandler_Handle_publishError_NackDiscard(t *testing.T) {
	ctrl := gomock.NewController(t)
	pub := publisher.NewMockPublisher(ctrl)
	queueName := "test queue"
	conf := config.QueueConfig{
		ScrapingResultQueue: queueName,
	}
	pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{queueName}), gomock.Any()).Times(1).
		Return(errors.New("publish error"))

	handl := NewApiSpecDocHandler(pub, conf)
	body := `{"FileUrl":"test url","IsNotifyUser":false}`
	delivery := rabbitmq.Delivery{
		Delivery: amqp091.Delivery{Body: []byte(body)},
	}
	action := handl.Handle(delivery)
	assert.Equal(t, rabbitmq.NackDiscard, action)
}

func TestApiSpecDocHandler_Handle_allCorrectNotificationFalse_called1TimeAck(t *testing.T) {
	ctrl := gomock.NewController(t)
	pub := publisher.NewMockPublisher(ctrl)
	queueName := "test queue"
	conf := config.QueueConfig{
		ScrapingResultQueue: queueName,
	}
	pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{queueName}), gomock.Any()).Times(1).Return(nil)

	handl := NewApiSpecDocHandler(pub, conf)
	body := `{"FileUrl":"test url","IsNotifyUser":false}`
	delivery := rabbitmq.Delivery{
		Delivery: amqp091.Delivery{Body: []byte(body)},
	}
	action := handl.Handle(delivery)
	assert.Equal(t, rabbitmq.Ack, action)
}

func TestApiSpecDocHandler_Handle_allCorrectNotificationFalse_called2TimesAck(t *testing.T) {
	ctrl := gomock.NewController(t)
	pub := publisher.NewMockPublisher(ctrl)
	resQName := "test queue"
	notQName := "test notification queue"
	conf := config.QueueConfig{
		ScrapingResultQueue: resQName,
		NotificationQueue:   notQName,
	}
	firstCall := pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{resQName}), gomock.Any()).Times(1).Return(nil)
	pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{notQName}), gomock.Any()).Times(1).Return(nil).After(firstCall)

	handl := NewApiSpecDocHandler(pub, conf)
	body := `{"FileUrl":"test url","IsNotifyUser":true}`
	delivery := rabbitmq.Delivery{
		Delivery: amqp091.Delivery{Body: []byte(body)},
	}
	action := handl.Handle(delivery)
	assert.Equal(t, rabbitmq.Ack, action)
}

func TestApiSpecDocHandler_Handle_notificationError_called2TimesAck(t *testing.T) {
	ctrl := gomock.NewController(t)
	pub := publisher.NewMockPublisher(ctrl)
	resQName := "test queue"
	notQName := "test notification queue"
	conf := config.QueueConfig{
		ScrapingResultQueue: resQName,
		NotificationQueue:   notQName,
	}
	firstCall := pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{resQName}), gomock.Any()).
		Times(1).
		Return(nil)
	pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{notQName}), gomock.Any()).Times(1).
		Return(errors.New("unexpected notification error")).
		After(firstCall)

	handl := NewApiSpecDocHandler(pub, conf)
	body := `{"FileUrl":"test url","IsNotifyUser":true}`
	delivery := rabbitmq.Delivery{
		Delivery: amqp091.Delivery{Body: []byte(body)},
	}
	action := handl.Handle(delivery)
	assert.Equal(t, rabbitmq.Ack, action)
}
