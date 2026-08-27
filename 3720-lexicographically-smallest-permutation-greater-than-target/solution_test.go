package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexGreaterPermutation(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		target   string
		expected string
	}{
		{"Example 1", "abc", "bba", "bca"},
		{"Example 2", "leet", "code", "eelt"},
		{"Example 3", "baba", "bbaa", ""},
		{"Example 400", "aa", "bb", ""},
		{"Example 419", "z", "a", "z"},
		{"Example 437", "z", "z", ""},
		{"Example 443", "aa", "aa", ""},
		{"Example 458", "ab", "ab", "ba"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := lexGreaterPermutation(test.s, test.target)
			assert.Equal(t, test.expected, actual)
		})
	}
}
