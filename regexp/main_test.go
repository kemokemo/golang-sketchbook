package main

import "testing"

func TestColumnToPrefix(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abcd0123", "ABCD-0123"},
		{"vwxy", "VWXY"},
	}
	for _, test := range tests {
		if got := columnToPrefix(test.input); got != test.expected {
			t.Errorf(`columnToPrefix(%q) = %q`, test.input, got)
		}
	}
}
