package leetcode

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	points := findCriticalPoints(head)
	if len(points) <= 1 {
		return []int{-1, -1}
	}
	if len(points) == 2 {
		diff := int(math.Abs(float64(points[0] - points[1])))
		return []int{diff, diff}
	}
	lastIndex := len(points) - 1
	max := points[lastIndex] - points[0]
	min := max

	for i := 0; i < len(points)-1; i++ {
		cur, nex := points[i], points[i+1]
		if nex-cur < min {
			min = nex - cur
		}
	}

	return []int{min, max}
}

func findCriticalPoints(head *ListNode) []int {
	canFind := true
	prev := head.Val
	if head.Next == nil {
		return make([]int, 0)
	}
	head = head.Next
	cur := head.Val
	if head.Next == nil {
		return make([]int, 0)
	}
	next := head.Next.Val
	index := 1
	points := make([]int, 0)

	for canFind {
		if isCritical(prev, cur, next) {
			points = append(points, index)
		}

		index++
		head = head.Next
		if head.Next == nil {
			canFind = false
			break
		}
		prev = cur
		cur = next
		next = head.Next.Val
	}
	return points
}

func isCritical(prev, cur, next int) bool {
	if prev < cur && cur > next {
		return true
	}
	if prev > cur && cur < next {
		return true
	}
	return false
}
