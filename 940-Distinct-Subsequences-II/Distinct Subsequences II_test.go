package leetcode

import (
	"strings"
	"testing"
)

const subseqMod = 1_000_000_007

// distinctSubseqIIReference is an independent implementation of the same
// problem, using the last-occurrence DP formulation instead of the
// character-bucket formulation in the solution:
//
//	dp[i]  = number of distinct subsequences (including the empty one) of s[:i]
//	dp[0]  = 1
//	dp[i]  = 2*dp[i-1]              (append s[i-1] to every subsequence of s[:i-1])
//	         - dp[last[c]-1]        (those already counted at the previous
//	                                 occurrence of the same character c)
//	answer = dp[n] - 1              (exclude the empty subsequence)
//
// All subtractions are guarded with +mod before reducing.
func distinctSubseqIIReference(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	var last [26]int
	for i := 1; i <= len(s); i++ {
		c := s[i-1] - 'a'
		dp[i] = (2 * dp[i-1]) % subseqMod
		if last[c] != 0 {
			dp[i] = (dp[i] - dp[last[c]-1] + subseqMod) % subseqMod
		}
		last[c] = i
	}
	return (dp[len(s)] - 1 + subseqMod) % subseqMod
}

// distinctSubseqIIBruteForce enumerates every non-empty subsequence of a short
// string (use only for n <= 12) and counts the distinct ones with a set.
func distinctSubseqIIBruteForce(s string) int {
	seen := make(map[string]struct{})
	for mask := 1; mask < 1<<len(s); mask++ {
		var b strings.Builder
		for i := 0; i < len(s); i++ {
			if mask&(1<<i) != 0 {
				b.WriteByte(s[i])
			}
		}
		seen[b.String()] = struct{}{}
	}
	return len(seen)
}

// cycleAlphabet returns the 26 lowercase letters repeated cyclically up to
// length n (e.g. n=3 -> "abc", n=27 -> "abc...z" + "a").
func cycleAlphabet(n int) string {
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < n; i++ {
		b.WriteByte('a' + byte(i%26))
	}
	return b.String()
}

// shortString truncates long inputs so failure messages stay readable.
func shortString(s string) string {
	if len(s) > 32 {
		return s[:16] + "..." + s[len(s)-16:]
	}
	return s
}

func TestDistinctSubseqII(t *testing.T) {
	// Length-2000 string cycling through all 26 letters. Expected value is
	// computed with the independent reference implementation:
	// distinctSubseqIIReference(...) == 995084322.
	mixed := cycleAlphabet(2000)

	var tests = []struct {
		name string
		s    string
		want int
	}{
		{"All distinct characters", "abc", 7},
		{"Duplicate middle character", "aba", 6},
		{"All same character", "aaa", 3},
		{"Single character", "z", 1},
		{"All same character length 4", "aaaa", 4},
		{"Alternating duplicates", "abab", 11},
		{"All 26 letters no repeats", "abcdefghijklmnopqrstuvwxyz", 67108863},
		{"Max length all same character", strings.Repeat("a", 2000), 2000},
		{"Max length mixed characters", mixed, distinctSubseqIIReference(mixed)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := distinctSubseqII(tt.s)
			if ans != tt.want {
				t.Errorf("distinctSubseqII(%q) = %d, want %d", shortString(tt.s), ans, tt.want)
			}
			// Cross-check the independent reference on the same input so both
			// formulations agree, not just with the precomputed expectation.
			if ref := distinctSubseqIIReference(tt.s); ref != ans {
				t.Errorf("reference distinctSubseqIIReference(%q) = %d, solution = %d", shortString(tt.s), ref, ans)
			}
		})
	}
}

// TestDistinctSubseqIIBruteForce verifies the solution against exhaustive
// set enumeration on short strings (n <= 8 keeps 2^n small).
func TestDistinctSubseqIIBruteForce(t *testing.T) {
	var tests = []struct {
		name string
		s    string
	}{
		{"Single character", "z"},
		{"Two distinct characters", "ab"},
		{"Two same characters", "aa"},
		{"Three distinct characters", "abc"},
		{"Duplicate middle character", "aba"},
		{"All same character", "aaaa"},
		{"Alternating duplicates", "abab"},
		{"Random short", "acbca"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := distinctSubseqIIBruteForce(tt.s)
			if got := distinctSubseqII(tt.s); got != want {
				t.Errorf("distinctSubseqII(%q) = %d, brute force = %d", tt.s, got, want)
			}
		})
	}
}
