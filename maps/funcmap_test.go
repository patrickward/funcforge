package maps_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/patrickward/funcforge/maps"
)

func TestNew(t *testing.T) {
	tests := []struct {
		pairs    []any
		expected map[string]any
		err      bool
	}{
		{[]any{"key", "value", "other", "value"}, map[string]any{"key": "value", "other": "value"}, false},
		{[]any{"a", 1, "b", 2}, map[string]any{"a": 1, "b": 2}, false},
		{[]any{"key", "value", "other"}, nil, true},
	}

	for _, tt := range tests {
		result, err := maps.FuncMap()["dict"].(func(...any) (map[string]any, error))(tt.pairs...)
		if tt.err {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		}
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		m        map[string]any
		key      string
		def      any
		expected any
	}{
		{map[string]any{"key": "value"}, "key", "default", "value"},
		{map[string]any{"key": "value"}, "missing", "default", "default"},
	}

	for _, tt := range tests {
		result := maps.FuncMap()["dictGet"].(func(map[string]any, string, any) any)(tt.m, tt.key, tt.def)
		assert.Equal(t, tt.expected, result)
	}
}

func TestKeys(t *testing.T) {
	tests := []struct {
		m        map[string]any
		expected []string
	}{
		{map[string]any{"b": 2, "a": 1}, []string{"a", "b"}},
		{map[string]any{"c": 3, "a": 1, "b": 2}, []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		result := maps.FuncMap()["dictKeys"].(func(map[string]any) []string)(tt.m)
		assert.Equal(t, tt.expected, result)
	}
}

func TestValues(t *testing.T) {
	tests := []struct {
		m        map[string]any
		expected []any
	}{
		{map[string]any{"b": 2, "a": 1}, []any{1, 2}},
		{map[string]any{"c": 3, "a": 1, "b": 2}, []any{1, 2, 3}},
	}

	for _, tt := range tests {
		result := maps.FuncMap()["dictValues"].(func(map[string]any) []any)(tt.m)
		assert.Equal(t, tt.expected, result)
	}
}

func TestPick(t *testing.T) {
	tests := []struct {
		m        map[string]any
		keys     []string
		expected map[string]any
	}{
		{map[string]any{"a": 1, "b": 2, "c": 3}, []string{"a", "c"}, map[string]any{"a": 1, "c": 3}},
		{map[string]any{"a": 1, "b": 2, "c": 3}, []string{"b"}, map[string]any{"b": 2}},
	}

	for _, tt := range tests {
		result := maps.FuncMap()["dictPick"].(func(map[string]any, ...string) map[string]any)(tt.m, tt.keys...)
		assert.Equal(t, tt.expected, result)
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		maps     []map[string]any
		expected map[string]any
	}{
		{[]map[string]any{{"a": 1, "b": 2}, {"b": 3, "c": 4}}, map[string]any{"a": 1, "b": 3, "c": 4}},
		{[]map[string]any{{"a": 1}, {"b": 2}, {"c": 3}}, map[string]any{"a": 1, "b": 2, "c": 3}},
	}

	for _, tt := range tests {
		result := maps.FuncMap()["dictMerge"].(func(...map[string]any) map[string]any)(tt.maps...)
		assert.Equal(t, tt.expected, result)
	}
}
