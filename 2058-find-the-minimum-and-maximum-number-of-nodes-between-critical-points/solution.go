package leetcode

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	min, max := findCriticalPoints(head)
	return []int{min, max}
}

func findCriticalPoints(head *ListNode) (int, int) {
	canFind := true
	prev := head.Val
	if head.Next == nil {
		return -1, -1
	}
	head = head.Next
	cur := head.Val
	if head.Next == nil {
		return -1, -1
	}
	next := head.Next.Val
	index := 1

	prevIndex := -1
	firstIndex := -1
	lastIndex := -1
	minDist := math.MaxInt
	for canFind {
		if isCritical(prev, cur, next) {
			if firstIndex == -1 {
				firstIndex = index
			} else {
				dist := index - prevIndex
				if dist < minDist {
					minDist = dist
				}
			}
			prevIndex = index
			lastIndex = index

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
	if minDist == math.MaxInt {
		return -1, -1
	}
	return minDist, lastIndex - firstIndex
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
