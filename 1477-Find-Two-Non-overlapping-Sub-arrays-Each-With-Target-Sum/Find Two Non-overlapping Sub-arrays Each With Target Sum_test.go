package leetcode

import (
	"testing"
)

func TestMinSumOfLengths(t *testing.T) {
	var tests = []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{
			name:   "Example 1 - two single-element subarrays",
			arr:    []int{3, 2, 2, 4, 3},
			target: 3,
			want:   2,
		},
		{
			name:   "Example 2 - multiple options, pick shortest total",
			arr:    []int{7, 3, 4, 7},
			target: 7,
			want:   2,
		},
		{
			name:   "Example 3 - only one subarray exists",
			arr:    []int{4, 3, 2, 6, 2, 3, 4},
			target: 6,
			want:   -1,
		},
		{
			name:   "No subarray at all",
			arr:    []int{1, 2, 3},
			target: 10,
			want:   -1,
		},
		{
			name:   "Single element",
			arr:    []int{5},
			target: 5,
			want:   -1,
		},
		{
			name:   "Two elements both equal target",
			arr:    []int{3, 3},
			target: 3,
			want:   2,
		},
		{
			name:   "Adjacent subarrays",
			arr:    []int{2, 3, 2, 3},
			target: 5,
			want:   4,
		},
		{
			name:   "Adjacent non-overlapping subarrays",
			arr:    []int{2, 2, 1, 3, 2},
			target: 4,
			want:   4,
		},
		{
			name:   "All elements positive, larger array",
			arr:    []int{1, 1, 1, 1, 1, 1, 1, 1},
			target: 2,
			want:   4,
		},
		{
			name:   "Target matches entire array value",
			arr:    []int{10, 20, 30},
			target: 10,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := minSumOfLengths(tt.arr, tt.target)
			if ans != tt.want {
				t.Errorf("minSumOfLengths(%v, %d) = %d, want %d",
					tt.arr, tt.target, ans, tt.want)
			}
		})
	}
}
