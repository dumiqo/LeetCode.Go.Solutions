package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCommasExamples(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		want int64
	}{
		{"n = 1 (no commas)", 1, 0},
		{"n = 998 (no commas)", 998, 0},
		{"n = 1000 (just 1,000)", 1000, 1},
		{"n = 1002 (1,000..1,002)", 1002, 3},
		{"n = 9999 (9,000 numbers each with one comma)", 9999, 9000},
		{"n = 10000 (plus 10,000)", 10000, 9001},
		{"n = 100000 (1000..100000)", 100000, 99001},
		{"n = 1000000 (crosses second comma)", 1000000, 999002},
		{"n = 1000000000 (crosses third comma)", 1000000000, 1998999003},
		{"n = 10^15 - 1 (just below max threshold)", 999999999999999, 3998998998999000},
		{"n = 1000000000000000 (max constraint)", 1000000000000000, 3998998998999005},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countCommas(tt.n))
		})
	}
}

// TestCountCommasAgainstBruteForce cross-checks the formula against a
// genuinely independent digit-length oracle over a range that crosses the
// first and second comma thresholds (10^3 and 10^6). Rather than recomputing
// a cumulative sum from scratch for every n (which would be O(n^2)), the
// oracle run is accumulated incrementally in O(n) total. The oracle counts
// a number's commas from its digit length alone ((digits-1)/3), with no
// shared threshold logic, so it independently validates the formula at every
// point.
func TestCountCommasAgainstBruteForce(t *testing.T) {
	const limit int64 = 1000000
	var cumulative int64
	for n := int64(1); n <= limit; n++ {
		cumulative += int64(commasByDigits(n))
		assert.Equal(t, cumulative, countCommas(n), "mismatch at n = %d", n)
	}
}

// commasByDigits returns the number of commas in the standard formatted
// representation of x using only its digit length: a d-digit number carries
// (d-1)/3 commas. This never builds a string, so it is a cheap independent
// oracle for a single number.
func commasByDigits(x int64) int {
	digits := 0
	for y := x; y > 0; y /= 10 {
		digits++
	}
	return (digits - 1) / 3
}
