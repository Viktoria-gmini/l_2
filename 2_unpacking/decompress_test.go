package main

import (
	"testing"
)

func TestDecompressString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a3b2c", "aaabbc"},
		{"x3w4z4", "xxxwwwwzzzz"},
		{"b5", "bbbbb"},
		{"abcd", "abcd"},
		{"de12", "deeeeeeeeeeee"},
		{"", ""},
		{"45", ""},
	}

	for _, test := range tests {
		result := DecompressString(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, but got %s for input %s", test.expected, result, test.input)
		}
	}
}
