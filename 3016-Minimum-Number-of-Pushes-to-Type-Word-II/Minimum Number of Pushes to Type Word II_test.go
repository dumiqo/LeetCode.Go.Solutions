package leetcode

import (
	"testing"
)

func TestMinimumPushes(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		// the table itself
		{"Should be 5", "abcde", 5},
		{"Should be 24", "aabbccddeeffgghhiiiiii", 24},
		{"Should be 12", "xyzxyzxyzxyz", 12},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := minimumPushes(tt.input)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
