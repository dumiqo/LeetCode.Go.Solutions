package leetcode

import (
	"testing"
)

func TestStoneGameII(t *testing.T) {
	var tests = []struct {
		name   string
		result int
		input  []int
	}{
		{"Alice max 10", 10, []int{2, 7, 9, 4, 4}},
		{"Alice max 104", 104, []int{1, 2, 3, 4, 5, 100}},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := stoneGameII(tt.input)
			if ans != tt.result {
				t.Errorf("got wrong result %v, want %v", ans, tt.result)
			}
		})
	}
}
