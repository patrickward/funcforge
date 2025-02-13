package debug

import (
	"encoding/json"
	"fmt"
	"html/template"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"dbgDump":   dump,   // pretty prints any value
		"dbgTypeof": typeof, // returns the type of a value
	}
}

// dump pretty prints any value
func dump(v any) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}

// typeof returns the type of a value
func typeof(v any) string {
	return fmt.Sprintf("%T", v)
}
