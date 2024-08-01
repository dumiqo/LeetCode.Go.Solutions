package leetcode

import (
	"testing"
)

func TestTwoSeniorCinizens(t *testing.T) {
	strArray := []string{
		"7868190130M7522",
		"5303914400F9211",
		"9273338290F4010",
	}

	result := countSeniors(strArray)
	if result != 2 {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, 2)
	}
}

func TestNoSeniorCitizens(t *testing.T) {
	strArray := []string{
		"1313579440F2036",
		"2921522980M5644",
	}

	result := countSeniors(strArray)
	if result != 0 {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, 0)
	}
}
