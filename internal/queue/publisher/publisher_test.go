package publisher

import (
	"github.com/golang/mock/gomock"
	mock_logger "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/logger/mocks"
	publisher "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/publisher/mocks"
	"testing"
)

func TestClosePublisher(t *testing.T) {
	ctrl := gomock.NewController(t)
	pub := publisher.NewMockPublisher(ctrl)
	log := mock_logger.NewMockLogger(ctrl)
	log.EXPECT().Info(gomock.Any())

	pub.EXPECT().Close().Return(nil)
	ClosePublisher(pub, log)
}
