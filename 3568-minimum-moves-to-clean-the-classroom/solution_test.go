package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMoves(t *testing.T) {
	tests := []struct {
		name      string
		classroom []string
		energy    int
		expected  int
	}{
		{"Test 1", []string{"S.", "XL"}, 2, 2},
		{"Test 2", []string{"LS", "RL"}, 4, 3},
		{"Test 3", []string{"L.S", "RXL"}, 3, -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := minMoves(test.classroom, test.energy)
			assert.Equal(t, test.expected, actual)
		})
	}
}
