package leetcode

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCommasExamples(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"n = 998 (no commas)", 998, 0},
		{"n = 1000 (just 1,000)", 1000, 1},
		{"n = 1002 (1,000..1,002)", 1002, 3},
		{"n = 9999 (9,000 numbers each with one comma)", 9999, 9000},
		{"n = 10000 (plus 10,000)", 10000, 9001},
		{"n = 100000 (1000..100000)", 100000, 99001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countCommas(tt.n))
		})
	}
}

func TestCountCommasAgainstBruteForce(t *testing.T) {
	for n := 1; n <= 10000; n++ {
		assert.Equal(t, bruteForceCountCommas(n), countCommas(n),
			"mismatch at n = %d", n)
	}
}

// bruteForceCountCommas sums commas by formatting each number individually.
func bruteForceCountCommas(n int) int {
	total := 0
	for x := 1; x <= n; x++ {
		total += commasInFormattedNumber(x)
	}
	return total
}

// commasInFormattedNumber returns the number of commas in the standard
// formatted representation of x (a comma after every three digits from the
// right).
func commasInFormattedNumber(x int) int {
	s := strconv.Itoa(x)
	var b strings.Builder
	b.Grow(len(s) + (len(s)-1)/3)
	for i := 0; i < len(s); i++ {
		b.WriteByte(s[i])
		remaining := len(s) - 1 - i
		if remaining > 0 && remaining%3 == 0 {
			b.WriteByte(',')
		}
	}
	return strings.Count(b.String(), ",")
}
