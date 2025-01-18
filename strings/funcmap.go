package strings

import (
	"fmt"
	"html/template"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caser = cases.Title(language.English)

// FuncMap returns a function map with functions for working with strings.
func FuncMap() template.FuncMap {
	return template.FuncMap{
		"strContains":  strings.Contains,  // Check if a string contains a substring
		"strHasPrefix": strings.HasPrefix, // Check if a string has a prefix
		"strHasSuffix": strings.HasSuffix, // Check if a string has a suffix
		"strJoin":      strings.Join,      // Join a slice into a string
		"strLower":     strings.ToLower,   // Convert a string to lowercase
		"strSplit":     strings.Split,     // Split a string into a slice
		"strTitleize":  Titleize,          // Capitalize the first letter of each word in a string
		"strToString":  ToString,          // Convert any type to a string
		"strTrim":      strings.TrimSpace, // Trim leading and trailing spaces from a string
		"strTrimAll":   strings.Trim,      // Trim a string using a set of characters
		"strTruncate":  Truncate,          // Truncate a string to a specified length
		"strUpper":     strings.ToUpper,   // Convert a string to uppercase
	}
}

// Titleize capitalizes the first letter of each word in a string (English)
func Titleize(s string) string {
	return caser.String(s)
}

// ToString converts any type to a string
func ToString(i any) string {
	return fmt.Sprintf("%v", i)
}

// Truncate truncates a string to a specified length and appends "..." if longer
func Truncate(length int, s string) string {
	if len(s) <= length {
		return s
	}
	// Find last space before length
	if idx := strings.LastIndex(s[:length], " "); idx != -1 {
		return s[:idx] + "..."
	}
	return s[:length] + "..."
}
