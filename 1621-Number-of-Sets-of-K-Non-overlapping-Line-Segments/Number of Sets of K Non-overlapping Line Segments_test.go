package leetcode

import (
	"testing"
)

func TestNumberOfSets(t *testing.T) {
	var tests = []struct {
		name string
		n    int
		k    int
		want int
	}{
		{"Example 1", 4, 2, 5},
		{"Example 2", 3, 1, 3},
		{"Example 3", 30, 7, 796297179},
		{"Minimum n", 2, 1, 1},
		{"All points segmented", 3, 2, 1},
		{"Five points three segments", 5, 3, 7},
		{"One segment on many points", 1000, 1, 499500},
		{"Maximum k", 1000, 999, 1},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := numberOfSets(tt.n, tt.k)
			if ans != tt.want {
				t.Errorf("numberOfSets(%d, %d) = %d, want %d", tt.n, tt.k, ans, tt.want)
			}
		})
	}
}
