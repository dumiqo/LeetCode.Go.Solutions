package leetcode

import (
	"testing"
)

func TestNumberToWords(t *testing.T) {
	var tests = []struct {
		name  string
		input int
		want  string
	}{
		// the table itself
		{"Parse 123", 123, "One Hundred Twenty Three"},
		{"Parse 1000000", 1000000, "One Million"},
		{"Parse 0", 0, "Zero"},
		{"Parse 12345", 12345, "Twelve Thousand Three Hundred Forty Five"},
		{"Parse 1234567", 1234567, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := numberToWords(tt.input)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
