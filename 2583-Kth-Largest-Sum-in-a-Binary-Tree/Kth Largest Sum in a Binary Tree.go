package leetcode

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	dict := map[int]int64{}
	build(dict, root, 0)

	if len(dict) < k {
		return -1
	}

	pl := make([]int64, len(dict))
	i := 0
	for _, v := range dict {
		pl[i] = v
		i++
	}
	sort.Slice(pl, func(i, j int) bool { return pl[i] > pl[j] })

	return pl[k-1]
}

func build(dict map[int]int64, node *TreeNode, depth int) {
	if node == nil {
		return
	}
	dict[depth] += int64(node.Val)
	depth += 1
	build(dict, node.Left, depth)
	build(dict, node.Right, depth)
}
