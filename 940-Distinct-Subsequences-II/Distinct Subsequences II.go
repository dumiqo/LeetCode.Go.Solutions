package leetcode

const mod = 1_000_000_007

// distinctSubseqII returns the number of distinct non-empty subsequences of s,
// modulo 1_000_000_007.
//
// Character-bucket DP: end[c] holds the number of distinct non-empty
// subsequences ending with character c (mod mod), and total is their sum.
// Appending character c to every existing subsequence and to the empty string
// creates total+1 new distinct subsequences ending with c; the previous count
// end[c] was already counted elsewhere and must be subtracted.
func distinctSubseqII(s string) int {
	var end [26]int
	total := 0
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		prev := end[c]
		end[c] = (total + 1) % mod
		total = (2*total + 1 - prev) % mod
		if total < 0 {
			total += mod
		}
	}
	return total
}
