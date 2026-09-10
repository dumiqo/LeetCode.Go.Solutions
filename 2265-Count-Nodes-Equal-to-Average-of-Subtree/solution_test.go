package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// rightSkewedTree builds a right-skewed chain of n nodes where node i (1-based
// from the root) has value i. The deepest node has value n.
func rightSkewedTree(n int) *TreeNode {
	var root *TreeNode
	for i := n; i >= 1; i-- {
		root = &TreeNode{Val: i, Right: root}
	}
	return root
}

// leftSkewedTree builds a left-skewed chain of n nodes where node i (1-based
// from the root) has value i. The deepest node has value n.
func leftSkewedTree(n int) *TreeNode {
	var root *TreeNode
	for i := n; i >= 1; i-- {
		root = &TreeNode{Val: i, Left: root}
	}
	return root
}

// leftSkewedTreeOf builds a left-skewed chain of n nodes, all with value v.
func leftSkewedTreeOf(n, v int) *TreeNode {
	var root *TreeNode
	for i := 0; i < n; i++ {
		root = &TreeNode{Val: v, Left: root}
	}
	return root
}

func TestAverageOfSubtree(t *testing.T) {
	tests := []struct {
		name   string
		root   *TreeNode
		expect int
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 1},
				},
				Right: &TreeNode{
					Val:   5,
					Right: &TreeNode{Val: 6},
				},
			},
			expect: 5,
		},
		{
			name:   "Example 2 single node",
			root:   &TreeNode{Val: 1},
			expect: 1,
		},
		{
			name:   "Single node with value zero",
			root:   &TreeNode{Val: 0},
			expect: 1,
		},
		{
			name:   "Root average does not match",
			root:   &TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
			expect: 2,
		},
		{
			name: "Every node matches its subtree average",
			root: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 5, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 5},
			},
			expect: 5,
		},
		{
			// Only the leaf (5) and the node 4 (avg of 4 and 5) match.
			name:   "Right-skewed tree",
			root:   rightSkewedTree(5),
			expect: 2,
		},
		{
			// Zero-valued node inside a multi-node subtree: the root average
			// (0 + 0 + 1) / 3 = 0 truncates toward zero and matches the root.
			name:   "Zero-valued node with truncating average",
			root:   &TreeNode{Val: 0, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 1}},
			expect: 3,
		},
		{
			name:   "Maximum constraint values",
			root:   &TreeNode{Val: 1000, Left: &TreeNode{Val: 1000}, Right: &TreeNode{Val: 1000}},
			expect: 3,
		},
		{
			// Left chain of 1000 nodes all valued 1000 (Node.Val max).
			name:   "Large tree at maximum value",
			root:   leftSkewedTreeOf(1000, 1000),
			expect: 1000,
		},
		{
			// Balanced tree: 2 with two leaf children 2, so every node matches.
			name:   "Balanced tree",
			root:   &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}},
			expect: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := averageOfSubtree(test.root)

			assert.Equal(t, test.expect, actual)
		})
	}
}

// TestAverageOfSubtreeDeepLeftChain exercises worst-case recursion depth
// (O(n) stack for a left-skewed tree with n = 1000, the constraint maximum).
func TestAverageOfSubtreeDeepLeftChain(t *testing.T) {
	// Values 1..1000 (root = 1, leaf = 1000). Only the two deepest nodes match:
	// node 1000 (1000 / 1) and node 999, whose subtree average is (999 + 1000) / 2
	// = 999 by integer division. Every node k <= 998 has subtree average
	// (k + 1000) / 2 > k.
	expect := 2
	actual := averageOfSubtree(leftSkewedTree(1000))

	assert.Equal(t, expect, actual)
}
