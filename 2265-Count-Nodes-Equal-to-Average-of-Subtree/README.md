# 2265. Count Nodes Equal to Average of Subtree

**Difficulty:** Medium

## Problem

Given the `root` of a binary tree, return the number of nodes where the value of the node is equal to the **average** of the values in its **subtree**.

- The **average** of `n` elements is the sum of the `n` elements divided by `n`, using **integer division**.
- The **subtree** of a node is the tree consisting of that node, plus the set of all descendants of that node.

### Example 1

```
Input: root = [4,8,5,0,1,null,6]
Output: 5
```

Explanation:

- For the node with value `4`: the average of its subtree is `(4 + 8 + 5 + 0 + 1 + 6) / 6 = 24 / 6 = 4`.
- For the node with value `5`: the average of its subtree is `(5 + 6) / 2 = 11 / 2 = 5`.
- For the node with value `0`: the average of its subtree is `0 / 1 = 0`.
- For the node with value `1`: the average of its subtree is `1 / 1 = 1`.
- For the node with value `6`: the average of its subtree is `6 / 1 = 6`.

### Example 2

```
Input: root = [1]
Output: 1
```

Explanation: for the node with value `1`, the average of its subtree is `1 / 1 = 1`.

### Constraints

- The number of nodes in the tree is in the range `[1, 1000]`.
- `0 <= Node.val <= 1000`

## Approach: Post-order DFS

A node's subtree average depends only on its left and right subtrees, so a single bottom-up (post-order) traversal is enough.

For every node we compute:

- `sum` — the sum of all values in the subtree rooted at this node (`node.Val + leftSum + rightSum`),
- `count` — the number of nodes in that subtree (`1 + leftCount + rightCount`).

If `sum / count == node.Val` (integer division), the node is counted.

We can accumulate the answer on the fly, so no extra per-node storage is needed.

- **Time complexity:** `O(n)`, where `n` is the number of nodes — each node is visited exactly once.
- **Space complexity:** `O(h)` for the recursion stack, where `h` is the height of the tree (`O(n)` in the worst case, `O(log n)` for a balanced tree).

## Go Solution

```go
func averageOfSubtree(root *TreeNode) int {
	var ans int
	postOrder(root, &ans)
	return ans
}

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
```

## Run Tests

```bash
go test ./2265-Count-Nodes-Equal-to-Average-of-Subtree/
```