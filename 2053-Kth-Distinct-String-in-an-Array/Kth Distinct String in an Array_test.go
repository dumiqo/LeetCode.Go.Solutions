package leetcode

import (
	"testing"
)

func TestSecondUniq(t *testing.T) {
	array := []string{"d", "b", "c", "b", "c", "a"}

	result := kthDistinct(array, 2)
	if result != "a" {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, "a")
	}
}

func TestFirstUniq(t *testing.T) {
	array := []string{"aaa", "aa", "a"}

	result := kthDistinct(array, 1)
	if result != "aaa" {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, "aaa")
	}
}

func TestEmptyString(t *testing.T) {
	array := []string{"a", "b", "a"}

	result := kthDistinct(array, 3)
	if result != "" {
		t.Errorf("Result was incorrect, got: %v, want: empty string.", result)
	}
}
