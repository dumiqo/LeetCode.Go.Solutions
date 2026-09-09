package leetcode

// countCommas returns the total number of commas used when writing all
// integers from 1 to n (inclusive) in standard number formatting, where a
// comma is inserted after every three digits from the right.
//
// A number x is written with at least k commas iff x >= 10^(3*k), because
// every additional comma requires three more leading digits (10^3 = 1,000;
// 10^6 = 1,000,000; ...). Summing, over k, the count of numbers in [1, n]
// that reach each threshold gives the total number of commas.
//
// n can be up to 10^15, so int64 is used throughout. The maximum answer
// (~4e15) fits comfortably in int64, so no modulo is needed. The overflow
// guard (pow > n/1000) prevents pow *= 1000 from exceeding the int64 range.
func countCommas(n int64) int64 {
	total := int64(0)
	for pow := int64(1000); pow <= n; {
		total += n - pow + 1
		if pow > n/1000 {
			break
		}
		pow *= 1000
	}
	return total
}
