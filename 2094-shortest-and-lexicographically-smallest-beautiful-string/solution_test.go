package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestBeautifulSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want string
	}{
		{"Example 1", "100011001", 3, "11001"},
		{"Example 2", "1011", 2, "11"},
		{"Example 3", "000", 1, ""},
		{"Example 636 ", "001110101101101111", 10, "10101101101111"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := shortestBeautifulSubstring(test.s, test.k)
			assert.Equal(t, test.want, got, "shortestBeautifulSubstring(%q)", test.s)
		})
	}
}
