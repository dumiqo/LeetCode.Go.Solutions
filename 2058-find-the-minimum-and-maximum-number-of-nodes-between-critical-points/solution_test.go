package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodesBetweenCriticalPoints(t *testing.T) {
	tests := []struct {
		name     string
		input    ListNode
		expected []int
	}{
		{"Example 1", ListNode{1, &ListNode{3, nil}}, []int{-1, -1}},
		{"Example 2", ListNode{5, &ListNode{3, &ListNode{1, &ListNode{2, &ListNode{5, &ListNode{1, &ListNode{2, nil}}}}}}}, []int{1, 3}},
		{"Example 3", ListNode{1, &ListNode{3, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{2, &ListNode{2, &ListNode{2, &ListNode{7, nil}}}}}}}}}, []int{3, 3}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := nodesBetweenCriticalPoints(&test.input)

			assert.Equal(t, test.expected, actual)
		})
	}
}
