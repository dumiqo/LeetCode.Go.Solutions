package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCommas(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		expect int
	}{
		{"Example 1", 1002, 3},
		{"Example 2", 998, 0},
		{"Minimum n", 1, 0},
		{"Just below first comma", 999, 0},
		{"First comma appears", 1000, 1},
		{"Just below second threshold", 9999, 9000},
		{"Second threshold reached", 10000, 9001},
		{"Just below third threshold", 99999, 99000},
		{"Maximum n", 100000, 99001},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := countCommas(test.n)

			assert.Equal(t, test.expect, actual)
		})
	}
}
