package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexPalindromicPermutation(t *testing.T) {
	tests := []struct {
		name, s, target, expected string
	}{
		{"Test Case 1", "baba", "abba", "baab"},
		{"Test Case 2", "baba", "bbaa", ""},
		{"Test Case 3", "abc", "abb", ""},
		{"Test Case 4", "aac", "aab", "aca"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := lexPalindromicPermutation(test.s, test.target)

			assert.Equal(t, test.expected, actual)
		})
	}
}
