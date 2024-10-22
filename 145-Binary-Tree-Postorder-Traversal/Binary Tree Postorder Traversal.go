package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	list := make([]int, 0)

	list = inner(root, list)
	return list
}

func inner(root *TreeNode, list []int) []int {
	if root == nil {
		return list
	}
	if root.Left == (&TreeNode{}) && root.Right == (&TreeNode{}) {
		return append(list, root.Val)
	}
	list = inner(root.Left, list)
	list = inner(root.Right, list)
	list = append(list, root.Val)
	return list
}
