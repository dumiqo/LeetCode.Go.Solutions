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
		{"Example 1", 1002, 3},            // "1,000", "1,001", "1,002" -> one comma each
		{"Example 2", 998, 0},             // all numbers have fewer than four digits
		{"number one", 1, 0},              // single digit, no comma
		{"up to 999", 999, 0},             // below 1000 no commas at all
		{"first comma", 1000, 1},          // "1,000"
		{"1999", 1999, 1000},              // 1000..1999, one comma each
		{"9999", 9999, 9000},              // 1000..9999 (9000 numbers, one comma each)
		{"10000", 10000, 9001},            // plus "10,000" (5 digits -> 1 comma)
		{"99999", 99999, 99000},           // 9000 (4-digit) + 90000 (5-digit)
		{"max constraint", 100000, 99001}, // plus "100,000" (6 digits -> 1 comma)
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := countCommas(test.n)
			assert.Equal(t, test.expect, actual)
		})
	}
}

// TestCountCommasBruteForce cross-checks countCommas against a naive per-number
// digit count over the full allowed input range.
func TestCountCommasBruteForce(t *testing.T) {
	for n := 1; n <= 100000; n++ {
		assert.Equal(t, bruteCountCommas(n), countCommas(n), "n=%d", n)
	}
}

// bruteCountCommas adds floor((digits-1)/3) per number in [1, n]: the number
// of commas needed to write it in standard US thousands-separator format.
func bruteCountCommas(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		digits := 0
		for x := i; x > 0; x /= 10 {
			digits++
		}
		total += (digits - 1) / 3
	}
	return total
}
