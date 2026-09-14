package leetcode

import (
	"reflect"
	"testing"
)

func TestMaximumWeight(t *testing.T) {
	var tests = []struct {
		name      string
		intervals [][]int
		want      []int
	}{
		{
			name:      "Example 1",
			intervals: [][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}, {6, 9, 3}, {6, 7, 1}, {8, 9, 1}},
			want:      []int{2, 3},
		},
		{
			name:      "Example 2",
			intervals: [][]int{{5, 8, 1}, {6, 7, 7}, {4, 7, 3}, {9, 10, 6}, {7, 8, 2}, {11, 14, 3}, {3, 5, 5}},
			want:      []int{1, 3, 5, 6},
		},
		{
			name:      "Single interval",
			intervals: [][]int{{1, 2, 10}},
			want:      []int{0},
		},
		{
			name:      "Two non-overlapping, take both",
			intervals: [][]int{{1, 2, 5}, {3, 4, 5}},
			want:      []int{0, 1},
		},
		{
			name:      "Overlapping tie-break lexicographically smaller",
			intervals: [][]int{{1, 4, 5}, {2, 3, 5}},
			want:      []int{0},
		},
		{
			name:      "Max 4 intervals from 5 non-overlapping identical",
			intervals: [][]int{{1, 2, 1}, {3, 4, 1}, {5, 6, 1}, {7, 8, 1}, {9, 10, 1}},
			want:      []int{0, 1, 2, 3},
		},
		{
			name:      "Share boundary -- overlapping (cannot take both)",
			intervals: [][]int{{1, 2, 1}, {2, 3, 9}},
			want:      []int{1},
		},
		{
			name:      "Equal weight, shorter array wins lexicographically",
			intervals: [][]int{{1, 10, 5}, {2, 3, 2}, {4, 5, 3}},
			want:      []int{0},
		},
		{
			name:      "All overlap, pick the heaviest",
			intervals: [][]int{{1, 10, 1}, {1, 10, 2}, {1, 10, 3}},
			want:      []int{2},
		},
		{
			name:      "Unsorted input, verify indexing correct",
			intervals: [][]int{{6, 7, 1}, {1, 3, 2}, {4, 5, 2}},
			want:      []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := maximumWeight(tt.intervals)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
