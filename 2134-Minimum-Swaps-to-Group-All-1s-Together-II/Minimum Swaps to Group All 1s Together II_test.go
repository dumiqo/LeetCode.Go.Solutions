package leetcode

import (
	"testing"
)

func TestOneSwipe(t *testing.T) {
	array := []int{0, 1, 0, 1, 1, 0, 0}

	result := minSwaps(array)
	if result != 1 {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, 1)
	}
}

func TestTwoSwipe(t *testing.T) {
	array := []int{0, 1, 1, 1, 0, 0, 1, 1, 0}

	result := minSwaps(array)
	if result != 2 {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, 2)
	}
}

func TestZeroSwipe(t *testing.T) {
	array := []int{1, 1, 0, 0, 1}

	result := minSwaps(array)
	if result != 0 {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, 0)
	}
}
