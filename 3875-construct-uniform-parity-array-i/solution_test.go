package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniformArray(t *testing.T) {
	tests := []struct {
		name   string
		nums1  []int
		expect bool
	}{
		{"Example 1", []int{2, 3}, true},
		{"Example 2", []int{4, 6}, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := uniformArray(test.nums1)

			assert.Equal(t, test.expect, actual)
		})
	}
}
