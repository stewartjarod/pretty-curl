package main

import "encoding/json"

// IsJSON returns true if the string is JSON
func IsJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
