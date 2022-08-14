package parse

import (
	"github.com/rog-golang-buddies/api-hub_data-scraping-service/internal/parse/openapi"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConverter(t *testing.T) {
	parsers := []Parser{openapi.NewYamlOpenApiParser(), openapi.NewJsonOpenApiParser()}
	converter := NewConverter(parsers)
	assert.NotNil(t, converter)
}
