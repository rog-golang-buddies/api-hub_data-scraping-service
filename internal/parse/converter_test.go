package parse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConverter(t *testing.T) {
	parsers := []Parser{NewYamlOpenApiParser(), NewJsonOpenApiParser()}
	converter := NewConverter(parsers)
	assert.NotNil(t, converter)
}
