package leetcode

import (
	"testing"
)

func TestNumberToWords(t *testing.T) {
	var tests = []struct {
		name  string
		left  *TreeNode
		right *TreeNode
		want  bool
	}{
		// the table itself
		{"Equals", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, true},
		{"OneDifferent", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, true},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := flipEquiv(tt.left, tt.right)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
