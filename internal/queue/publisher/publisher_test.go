package publisher

import (
	"github.com/golang/mock/gomock"
	publisher "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/queue/publisher/mocks"
	"testing"
)

func TestClosePublisher(t *testing.T) {
	ctrl := gomock.NewController(t)
	pub := publisher.NewMockPublisher(ctrl)
	pub.EXPECT().Close().Return(nil)
	ClosePublisher(pub)
}
