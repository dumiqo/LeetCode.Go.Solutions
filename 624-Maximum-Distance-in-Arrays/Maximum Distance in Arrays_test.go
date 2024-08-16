package leetcode

import (
	"testing"
)

func TestMaxDistance(t *testing.T) {
	var tests = []struct {
		name   string
		input  [][]int
		result int
	}{
		{"Max distance = 4, three arrays", [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}, 4},
		{"One array input", [][]int{{1, 1}}, 0},
		{"Max distance = 4, two arrays", [][]int{{1, 4}, {0, 5}}, 4},
		{"Max distance = 5, two arrays", [][]int{{1, 4, 5}, {0, 2}}, 5},
		{"Max distance = 3, two arrays", [][]int{{-3, -2, 1}, {-2}}, 3},
		{"Max distance = 1, two arrays", [][]int{{1, 2, 2}, {1, 1, 2}}, 1},
		{"Max distance = 6, three arrays", [][]int{{-1, 1}, {-3, 1, 4}, {-2, -1, 0, 2}}, 6},
		{"Max distance = 8, three arrays", [][]int{{-2, 1, 4, 5}, {-2, 5, 6}, {-2, 0}}, 8},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := maxDistance(tt.input)
			if ans != tt.result {
				t.Errorf("got wrong result %v, want %v", ans, tt.result)
			}
		})
	}
}
