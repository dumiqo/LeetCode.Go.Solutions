package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumGame(t *testing.T) {
	tests := []struct {
		name string
		num  string
		want bool
	}{
		{"пример из условия 1", "5023", false},
		{"пример из условия 2", "25??", true},
		{"пример из условия 3", "?3295???", false},
		{"все вопросы", "????", true},
		{"без вопросов, разные суммы", "1234", true},
		{"без вопросов, равные суммы", "1111", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sumGame(test.num)
			assert.Equal(t, test.want, got, "sumGame(%q)", test.num)
		})
	}
}
