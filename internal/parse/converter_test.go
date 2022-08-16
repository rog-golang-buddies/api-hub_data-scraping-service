package parse

import (
	"github.com/golang/mock/gomock"
	mock_logger "github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/logger/mocks"
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/parse/openapi"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConverter(t *testing.T) {
	ctrl := gomock.NewController(t)
	log := mock_logger.NewMockLogger(ctrl)
	parsers := []Parser{openapi.NewOpenApi(log)}
	converter := NewConverter(parsers)
	assert.NotNil(t, converter)
}
