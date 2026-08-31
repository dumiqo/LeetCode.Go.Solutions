package leetcode

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	if head == nil {
		return []int{-1, -1}
	}
	minDist, maxDist := findCriticalPoints(head)
	return []int{minDist, maxDist}
}

func findCriticalPoints(head *ListNode) (int, int) {
	if head.Next == nil || head.Next.Next == nil {
		return -1, -1
	}

	prev := head
	cur := head.Next
	next := head.Next.Next
	index := 1 // индекс текущего узла (cur)

	firstIdx := -1
	lastIdx := -1
	prevIdx := -1
	minDist := math.MaxInt

	for next != nil {
		if (prev.Val < cur.Val && cur.Val > next.Val) ||
			(prev.Val > cur.Val && cur.Val < next.Val) {
			if firstIdx == -1 {
				firstIdx = index
			} else {
				dist := index - prevIdx
				if dist < minDist {
					minDist = dist
				}
			}
			prevIdx = index
			lastIdx = index
		}

		prev, cur, next = cur, next, next.Next
		index++
	}

	if firstIdx == -1 || lastIdx == firstIdx {
		return -1, -1
	}
	return minDist, lastIdx - firstIdx
}
