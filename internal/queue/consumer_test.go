package queue_test

import (
	"github.com/golang/mock/gomock"
	mock_logger "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/logger/mocks"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue"
	mock_queue "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/mocks"
	"testing"
)

func TestClosePublisher(t *testing.T) {
	ctrl := gomock.NewController(t)
	consumer := mock_queue.NewMockConsumer(ctrl)
	consumer.EXPECT().Close().Return(nil)
	log := mock_logger.NewMockLogger(ctrl)
	log.EXPECT().Info(gomock.Any())
	queue.CloseConsumer(consumer, log)
}
