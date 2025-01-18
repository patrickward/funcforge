package values_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/patrickward/funcforge/values"
)

func TestYesNo(t *testing.T) {
	tests := []struct {
		input    bool
		expected string
	}{
		{true, "Yes"},
		{false, "No"},
	}

	for _, tt := range tests {
		result := values.FuncMap()["valYesNo"].(func(bool) string)(tt.input)
		assert.Equal(t, tt.expected, result)
	}
}

func TestOnOff(t *testing.T) {
	tests := []struct {
		input    bool
		expected string
	}{
		{true, "On"},
		{false, "Off"},
	}

	for _, tt := range tests {
		result := values.FuncMap()["valOnOff"].(func(bool) string)(tt.input)
		assert.Equal(t, tt.expected, result)
	}
}
