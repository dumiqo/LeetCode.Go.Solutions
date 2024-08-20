package leetcode

import (
	"testing"
)

func TestStoneGame(t *testing.T) {
	var tests = []struct {
		name   string
		result bool
		input  []int
	}{
		{"Alice won 1", true, []int{5, 3, 4, 5}},
		{"Alice won 2", true, []int{3, 7, 2, 3}},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := stoneGame(tt.input)
			if ans != tt.result {
				t.Errorf("got wrong result %v, want %v", ans, tt.result)
			}
		})
	}
}
