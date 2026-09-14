package leetcode

import (
	"testing"
)

func TestIsRectangleOverlap(t *testing.T) {
	var tests = []struct {
		name string
		rec1 []int
		rec2 []int
		want bool
	}{
		{"Overlapping rectangles", []int{0, 0, 2, 2}, []int{1, 1, 3, 3}, true},
		{"Touch at an edge", []int{0, 0, 1, 1}, []int{1, 0, 2, 1}, false},
		{"Identical rectangles", []int{0, 0, 1, 1}, []int{0, 0, 1, 1}, true},
		{"Touch at a corner", []int{0, 0, 1, 1}, []int{1, 1, 2, 2}, false},
		{"Containment", []int{0, 0, 3, 3}, []int{1, 1, 2, 2}, true},
		{"Fully separated", []int{0, 0, 1, 1}, []int{2, 2, 3, 3}, false},
		{"Overlap with negative coordinates", []int{-3, -1, 1, 2}, []int{-1, -2, 0, 1}, true},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := isRectangleOverlap(tt.rec1, tt.rec2)
			if ans != tt.want {
				t.Errorf("got wrong result %v, want %v", ans, tt.want)
			}
		})
	}
}
