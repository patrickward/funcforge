package maps

import (
	"fmt"
	"html/template"
	"sort"
)

// FuncMap returns all map-related template functions
func FuncMap() template.FuncMap {
	return template.FuncMap{
		"dict":       newMap, // Was: "dict" in DefaultFuncMap
		"dictGet":    get,    // N/A - New function
		"dictSet":    set,    // N/A - New function
		"dictKeys":   keys,   // N/A - New function
		"dictValues": values, // N/A - New function
		"dictPick":   pick,   // N/A - New function
		"dictMerge":  merge,  // N/A - New function
	}
}

// newMap creates a new map from key-value pairs
//
// Example: {{ map.new "key" "value" "other" "value" }} -> map[key:value other:value]
func newMap(pairs ...any) (map[string]any, error) {
	if len(pairs)%2 != 0 {
		return nil, fmt.Errorf("map.new requires pairs of arguments")
	}

	result := make(map[string]any)
	for i := 0; i < len(pairs); i += 2 {
		key, ok := pairs[i].(string)
		if !ok {
			return nil, fmt.Errorf("map key must be string, got %T", pairs[i])
		}
		result[key] = pairs[i+1]
	}
	return result, nil
}

// get safely gets a value with default
//
// Example: {{ map.get .Map "key" "default" }} -> value or default
func get(m map[string]any, key string, def any) any {
	if val, ok := m[key]; ok {
		return val
	}
	return def
}

// set sets a value in a map
//
// Example: {{ map.set .Map "key" "value" }} -> map[key:value]
func set(m map[string]any, key string, value any) map[string]any {
	m[key] = value
	return m
}

// keys returns sorted keys of map
//
// Example: {{ map.keys .Map }} -> [key1 key2]
func keys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// values returns values of map in key order
//
// Example: {{ map.values .Map }} -> [value1 value2]
func values(m map[string]any) []any {
	keys := keys(m)
	vals := make([]any, len(keys))
	for i, k := range keys {
		vals[i] = m[k]
	}
	return vals
}

// pick returns new map with only specified keys
//
// Example: {{ map.pick .Map "key1" "key2" }} -> map[key1:value1 key2:value2]
func pick(m map[string]any, keys ...string) map[string]any {
	result := make(map[string]any)
	for _, k := range keys {
		if v, ok := m[k]; ok {
			result[k] = v
		}
	}
	return result
}

// merge combines maps, later values override earlier ones
//
// Example: {{ map.merge .Map1 .Map2 }} -> map[key1:value1 key2:value2 key3:value3]
func merge(maps ...map[string]any) map[string]any {
	result := make(map[string]any)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
