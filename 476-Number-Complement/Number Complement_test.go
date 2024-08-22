package leetcode

import (
	"testing"
)

func TestFindComplement(t *testing.T) {
	var tests = []struct {
		name   string
		result int
		input  int
	}{
		{"Complement for 5", 2, 5},
		{"Complement for 1", 0, 1},
		{"Complement for 3", 0, 3},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := findComplement(tt.input)
			if ans != tt.result {
				t.Errorf("got wrong result %v, want %v", ans, tt.result)
			}
		})
	}
}
