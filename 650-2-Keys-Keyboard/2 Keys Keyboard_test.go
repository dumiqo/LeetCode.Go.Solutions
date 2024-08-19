package leetcode

import (
	"testing"
)

func TestMinSteps(t *testing.T) {
	var tests = []struct {
		name   string
		result int
		input  int
	}{
		{"Distance for 3", 3, 3},
		{"Distance for 1", 0, 1},
		{"Distance for 1000", 21, 1000},
		{"Distance for 500", 19, 500},
		{"Distance for 502", 253, 502},
		{"Distance for 503", 503, 503},
		{"Distance for 501", 170, 501},
		{"Distance for 50", 12, 50},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := minSteps(tt.input)
			if ans != tt.result {
				t.Errorf("got wrong result %v, want %v", ans, tt.result)
			}
		})
	}
}
