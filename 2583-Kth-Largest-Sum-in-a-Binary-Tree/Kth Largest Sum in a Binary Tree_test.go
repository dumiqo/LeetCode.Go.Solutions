package leetcode

import (
	"testing"
)

func TestNumberToWords(t *testing.T) {
	var tests = []struct {
		name  string
		tree  *TreeNode
		index int
		want  int64
	}{
		// the table itself
		{"First in order", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, 1, 3},
		{"Second in order", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, 2, 2},
		{"Last in order", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, 3, 1},
		{"Over order", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, 4, -1},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := kthLargestLevelSum(tt.tree, tt.index)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
