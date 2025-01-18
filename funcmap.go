package funcmap

import (
	"html/template"

	"github.com/patrickward/funcforge/attr"
	"github.com/patrickward/funcforge/collections"
	"github.com/patrickward/funcforge/conversions"
	"github.com/patrickward/funcforge/core"
	"github.com/patrickward/funcforge/debug"
	"github.com/patrickward/funcforge/html"
	"github.com/patrickward/funcforge/maps"
	"github.com/patrickward/funcforge/numbers"
	"github.com/patrickward/funcforge/slices"
	"github.com/patrickward/funcforge/strings"
	"github.com/patrickward/funcforge/time"
	"github.com/patrickward/funcforge/url"
	"github.com/patrickward/funcforge/values"
)

// MergeFuncMaps merges the provided function maps into a single function map.
func MergeFuncMaps(maps ...template.FuncMap) template.FuncMap {
	result := make(template.FuncMap)
	for _, m := range maps {
		for key, value := range m {
			result[key] = value
		}
	}
	return result
}

// MergeIntoFuncMap merges the provided function maps into the provided function map.
func MergeIntoFuncMap(dst template.FuncMap, maps ...template.FuncMap) {
	for _, src := range maps {
		for key, value := range src {
			dst[key] = value
		}
	}
}

// cachedFuncMap holds the cached function map for the templates package.
var cachedFuncMap template.FuncMap

// FuncMap returns the complete function map for the templates package.
func FuncMap() template.FuncMap {
	if cachedFuncMap != nil {
		return cachedFuncMap
	}

	cachedFuncMap = MergeFuncMaps(
		core.FuncMap(),
		attr.FuncMap(),
		collections.FuncMap(),
		conversions.FuncMap(),
		debug.FuncMap(),
		html.FuncMap(),
		maps.FuncMap(),
		numbers.FuncMap(),
		slices.FuncMap(),
		strings.FuncMap(),
		time.FuncMap(),
		url.FuncMap(),
		values.FuncMap(),
	)

	return cachedFuncMap
}
