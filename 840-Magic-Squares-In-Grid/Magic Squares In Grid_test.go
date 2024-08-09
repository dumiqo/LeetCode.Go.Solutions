package leetcode

import "testing"

func TestNumMagicSquaresInside(t *testing.T) {
	var tests = []struct {
		name   string
		input  [][]int
		result int
	}{
		{"One magic square", [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}, 1},
		{"One element array input", [][]int{{8}}, 0},
		{"Zero magic square", [][]int{{7, 0, 5}, {2, 4, 6}, {3, 8, 1}}, 0},
		{"One magic square, in big array", [][]int{{3, 2, 9, 2, 7}, {6, 1, 8, 4, 2}, {7, 5, 3, 2, 7}, {2, 9, 4, 9, 6}, {4, 3, 8, 2, 5}}, 1},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := numMagicSquaresInside(tt.input)
			if ans != tt.result {
				t.Errorf("got wrong result %v, want %v", ans, tt.result)
			}
		})
	}
}
