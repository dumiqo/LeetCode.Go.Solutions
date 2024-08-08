package leetcode

import (
	"testing"
)

func TestSpiralMatrixIII(t *testing.T) {
	var tests = []struct {
		name   string
		rows   int
		cols   int
		rStart int
		cStart int
		result [][]int
		len    int
	}{
		{"One Col Table", 1, 4, 0, 0, [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}, 4},
		{"Table", 5, 6, 1, 4, [][]int{{1, 4}, {1, 5}, {2, 5}, {2, 4}, {2, 3}, {1, 3}, {0, 3}, {0, 4}, {0, 5}, {3, 5}, {3, 4}, {3, 3}}, 30},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := spiralMatrixIII(tt.rows, tt.cols, tt.rStart, tt.cStart)
			if len(ans) != tt.len {
				t.Errorf("got wrong length %v, want %v", len(ans), tt.len)
			}
			for i, v := range tt.result {
				var coord = ans[i]
				if coord[0] != v[0] {
					t.Errorf("got wrong col %v, need %v. index: %v", coord[0], v[0], i)
				}
				if coord[1] != v[1] {
					t.Errorf("got wrong row %v, need %v. index: %v", coord[1], v[1], i)
				}
			}
		})
	}
}
