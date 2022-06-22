package helper

import (
	"bytes"
	"strings"
)

// OneIsEmpty test if one of the pointers is nil or reference an empty string
func OneIsEmpty(values ...*string) bool {
	for _, v := range values {
		if v == nil || len(strings.TrimSpace(*v)) == 0 {
			return true
		}
	}
	return false
}

// Truncate string
func Truncate(value string, size int) string {
	runes := bytes.Runes([]byte(value))
	if len(runes) > size {
		return string(runes[:size]) + "..."
	}
	return string(runes)
}
