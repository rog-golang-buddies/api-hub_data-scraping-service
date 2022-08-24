package handler

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/config"
	mock_logger "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/logger/mocks"
	process "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/process/mocks"
	publisher "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/publisher/mocks"
	"github.com/rog-golang-buddies/api_hub_common/apispecdoc"
	"github.com/stretchr/testify/assert"
	"github.com/wagslane/go-rabbitmq"
	"testing"
)

func TestApiSpecDocHandler_Handle_wrongBody_NackDiscard(t *testing.T) {
	ctrl := gomock.NewController(t)
	pub := publisher.NewMockPublisher(ctrl)
	proc := process.NewMockUrlProcessor(ctrl)
	log := mock_logger.NewMockLogger(ctrl)
	log.EXPECT().Infof(gomock.Any(), gomock.Any())
	log.EXPECT().Errorf(gomock.Any(), gomock.Any())
	conf := config.QueueConfig{}

	handl := NewApiSpecDocHandler(pub, conf, proc, log)
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
	log := mock_logger.NewMockLogger(ctrl)
	log.EXPECT().Infof(gomock.Any(), gomock.Any())
	log.EXPECT().Error(gomock.Any())
	queueName := "test queue"
	conf := config.QueueConfig{
		ScrapingResultQueue: queueName,
	}
	proc.EXPECT().Process(gomock.Any(), "test url").Times(1).Return(&apispecdoc.ApiSpecDoc{}, nil)
	pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{queueName}), gomock.Any()).Times(1).
		Return(errors.New("publish error"))

	handl := NewApiSpecDocHandler(pub, conf, proc, log)
	body := `{"file_url":"test url","is_notify_user":false}`
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
	log := mock_logger.NewMockLogger(ctrl)
	log.EXPECT().Infof(gomock.Any(), gomock.Any())
	log.EXPECT().Info(gomock.Any())
	queueName := "test queue"
	conf := config.QueueConfig{
		ScrapingResultQueue: queueName,
	}
	proc.EXPECT().Process(gomock.Any(), "test url").Times(1).Return(&apispecdoc.ApiSpecDoc{}, nil)
	pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{queueName}), gomock.Any()).Times(1).Return(nil)

	handl := NewApiSpecDocHandler(pub, conf, proc, log)
	body := `{"file_url":"test url","is_notify_user":false}`
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
	log := mock_logger.NewMockLogger(ctrl)
	log.EXPECT().Infof(gomock.Any(), gomock.Any())
	log.EXPECT().Info(gomock.Any())
	resQName := "test queue"
	notQName := "test notification queue"
	conf := config.QueueConfig{
		ScrapingResultQueue: resQName,
		NotificationQueue:   notQName,
	}
	proc.EXPECT().Process(gomock.Any(), "test url").Times(1).Return(&apispecdoc.ApiSpecDoc{}, nil)
	firstCall := pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{resQName}), gomock.Any()).Times(1).Return(nil)
	pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{notQName}), gomock.Any()).Times(1).Return(nil).After(firstCall)

	handl := NewApiSpecDocHandler(pub, conf, proc, log)
	body := `{"file_url":"test url","is_notify_user":true}`
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
	log := mock_logger.NewMockLogger(ctrl)
	log.EXPECT().Infof(gomock.Any(), gomock.Any())
	log.EXPECT().Info(gomock.Any())
	log.EXPECT().Error(gomock.Any())
	resQName := "test queue"
	notQName := "test notification queue"
	conf := config.QueueConfig{
		ScrapingResultQueue: resQName,
		NotificationQueue:   notQName,
	}
	proc.EXPECT().Process(gomock.Any(), "test url").Times(1).Return(&apispecdoc.ApiSpecDoc{}, nil)
	firstCall := pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{resQName}), gomock.Any()).
		Times(1).
		Return(nil)
	pub.EXPECT().Publish(gomock.Any(), gomock.Eq([]string{notQName}), gomock.Any()).Times(1).
		Return(errors.New("unexpected notification error")).
		After(firstCall)

	handl := NewApiSpecDocHandler(pub, conf, proc, log)
	body := `{"file_url":"test url","is_notify_user":true}`
	delivery := rabbitmq.Delivery{
		Delivery: amqp091.Delivery{Body: []byte(body)},
	}
	action := handl.Handle(context.Background(), delivery)
	assert.Equal(t, rabbitmq.Ack, action)
}
