package tools

import (
	"bytes"
	"encoding/json"
)

// KeyInMap checks if a string key is in a map
func KeyInMap(key string, value map[string]any) bool {
	for k := range value {
		if k == key {
			return true
		}
	}
	return false
}

// ListContains checks if string is in list
func ListContains(list []string, value string) bool {
	for _, elem := range list {
		if elem == value {
			return true
		}
	}
	return false
}

// stringify tries to return a string version of value.
//
// If value is not a string it will try to json marshal the value and return the json string.
// If it can not be json marshaled it will return the input value.
//
// This is most useful when logging as log masking can only be done on strings.
// It would have been better if it could be done in a library level to reduce overhead,
// but tflog does not support this.
func stringify(value any) any {

	if valueStr, ok := value.(string); ok {
		return valueStr
	}

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(value)
	if err == nil {
		return buffer.String()
	}
	return value
}
