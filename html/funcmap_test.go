package html_test

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/patrickward/funcforge/html"
)

func TestSafeHTML(t *testing.T) {
	tests := []struct {
		input    string
		expected template.HTML
	}{
		{"<div>Test</div>", template.HTML("<div>Test</div>")},
		{"<script>alert('xss')</script>", template.HTML("<script>alert('xss')</script>")},
	}

	for _, tt := range tests {
		result := html.FuncMap()["htmlSafe"].(func(string) template.HTML)(tt.input)
		assert.Equal(t, tt.expected, result)
	}
}
