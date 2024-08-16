package leetcode

import (
	"testing"
)

func TestДemonadeChange(t *testing.T) {
	var tests = []struct {
		name   string
		input  []int
		result bool
	}{
		{"Enought money for change", []int{5, 5, 5, 10, 20}, true},
		{"Has no change", []int{5, 5, 10, 10, 20}, false},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := lemonadeChange(tt.input)
			if ans != tt.result {
				t.Errorf("got wrong result %v, want %v", ans, tt.result)
			}
		})
	}
}
