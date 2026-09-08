package leetcode

// countCommas returns the total number of commas used when writing all
// integers from 1 to n (inclusive) in standard number formatting, where a
// comma is inserted after every three digits from the right and numbers with
// fewer than four digits contain no commas.
//
// A number with d digits is written with (d-1)/3 commas. For each digit
// length d, the d-digit numbers not exceeding n form the interval
// [10^(d-1), min(n, 10^d-1)] and each contributes (d-1)/3 commas, so the
// answer is the sum over all digit lengths of
//
//	count * (d-1)/3,  where count = max(0, min(n, 10^d-1) - 10^(d-1) + 1).
//
// Values n <= 10^5 span at most six digit lengths, giving O(log n) time and
// O(1) space. Example: n = 1002 -> "1,000", "1,001", "1,002" each carry one
// comma, so the result is 3.
func countCommas(n int) int {
	total := 0
	// d-digit numbers occupy [low, high] = [10^(d-1), 10^d-1].
	for d, low, high := 1, 1, 9; low <= n; d, low, high = d+1, high+1, high*10+9 {
		end := n
		if high < end {
			end = high
		}
		count := end - low + 1
		total += count * ((d - 1) / 3)
	}
	return total
}
