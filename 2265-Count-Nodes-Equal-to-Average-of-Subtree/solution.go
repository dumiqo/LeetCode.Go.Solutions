package leetcode

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// averageOfSubtree returns the number of nodes whose value is equal to the
// integer average of the values in their subtree.
func averageOfSubtree(root *TreeNode) int {
	var ans int
	postOrder(root, &ans)
	return ans
}

// postOrder traverses the tree bottom-up. It returns the sum of node values
// and the number of nodes of the subtree rooted at node.
func postOrder(node *TreeNode, ans *int) (sum, count int) {
	if node == nil {
		return 0, 0
	}

	leftSum, leftCount := postOrder(node.Left, ans)
	rightSum, rightCount := postOrder(node.Right, ans)

	sum = node.Val + leftSum + rightSum
	count = 1 + leftCount + rightCount

	if sum/count == node.Val {
		*ans++
	}

	return sum, count
}
