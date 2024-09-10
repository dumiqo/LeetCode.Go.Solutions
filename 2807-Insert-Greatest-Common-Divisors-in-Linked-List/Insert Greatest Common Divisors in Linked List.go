package leetcode

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	var first = head

	for head.Next != nil {
		gcd := ListNode{Val: GCD(head.Val, head.Next.Val), Next: head.Next}
		head.Next = &gcd
		head = gcd.Next
	}

	return first
}

func GCD(a int, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

type ListNode struct {
	Val  int
	Next *ListNode
}
