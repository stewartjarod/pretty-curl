package main

import (
	"bytes"
	"encoding/json"
	"os"
)

// IsJSON returns true if the string is JSON
func IsJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

// PrintJSON outputs json to stdout
func PrintJSON(b []byte) {
	var out bytes.Buffer
	json.Indent(&out, b, "", "    ")
	out.WriteTo(os.Stdout)
}
