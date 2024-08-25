package leetcode

import (
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	var tests = []struct {
		name   string
		result []int
		input  TreeNode
	}{
		{"Treverse 1 null 2 3 tree", []int{3, 2, 1}, TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := postorderTraversal(tt.input)
			if len(ans) != len(tt.result) {
				t.Errorf("got wrong result length %v, want %v", len(ans), len(tt.result))
			}

			for i, tn := range ans {
				if tn != tt.result[i] {
					t.Errorf("got wrong result %v, want %v, id %v", tn, tt.result[i], i)
				}
			}
		})
	}
}
