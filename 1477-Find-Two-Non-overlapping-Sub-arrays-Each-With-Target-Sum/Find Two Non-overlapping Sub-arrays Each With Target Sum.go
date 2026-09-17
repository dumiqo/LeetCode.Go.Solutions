package leetcode

// minSumOfLengths returns the minimum sum of the lengths of two
// non-overlapping sub-arrays of arr whose sums both equal target.
// It returns -1 if no such pair of sub-arrays exists.
//
// A hash table records the last index at which each prefix sum occurred.
// f[i] keeps the minimum length of a target-sum sub-array found within the
// first i elements (f[i] is monotonic in i). When a target-sum sub-array ends
// at position i, its start is j = d[s-target], so its length is i-j. Combining
// it with the best sub-array fully to its left (which is f[j]) yields a
// candidate answer f[j] + (i - j). Since j < i, the two sub-arrays cannot
// overlap.
//
// Time:  O(n)
// Space: O(n)
func minSumOfLengths(arr []int, target int) int {
	const inf = 1 << 30

	d := map[int]int{0: 0} // prefix sum -> last index (1-based position)
	s, n := 0, len(arr)
	f := make([]int, n+1)
	f[0] = inf
	ans := inf

	for i, v := range arr {
		i++ // convert to 1-based position
		s += v
		f[i] = f[i-1]
		if j, ok := d[s-target]; ok {
			f[i] = min(f[i], i-j)
			ans = min(ans, f[j]+i-j)
		}
		d[s] = i
	}

	if ans > n {
		return -1
	}
	return ans
}
