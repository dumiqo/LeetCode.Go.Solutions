package leetcode

// countCommas returns the total number of commas used when writing all
// integers from 1 to n (inclusive) in standard number formatting, where a
// comma is inserted after every three digits from the right.
//
// A number x is written with at least k commas iff x >= 10^(3*k), because
// every additional comma requires three more leading digits. Summing, over
// k, the count of numbers in [1, n] that reach each threshold gives the
// total number of commas.
func countCommas(n int) int {
	total := 0
	for pow := 1000; pow <= n; {
		total += n - pow + 1
		if pow > n/1000 {
			break
		}
		pow *= 1000
	}
	return total
}
