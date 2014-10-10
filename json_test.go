package main

import "testing"

var testCases = []struct {
	jsonString string
	expected   bool
}{
	{"{ \"JSON\": 1, \"Test\": true}", true},
	{"{ \"JSON\": 1, \"Test\": true,}", false},
	{"{ \"JSON\": 1 }", true},
	{"{ JSON: 1 }", false},
	{"2400", false},
}

func TestJSON(t *testing.T) {
	for _, test := range testCases {
		observed := IsJSON(test.jsonString)

		if observed != test.expected {
			t.Fatalf("%s returned %v but should be %v", test.jsonString, observed, test.expected)
		}
	}
}
