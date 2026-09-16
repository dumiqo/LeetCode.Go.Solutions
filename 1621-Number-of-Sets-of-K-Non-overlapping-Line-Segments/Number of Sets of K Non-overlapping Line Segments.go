package leetcode

// numberOfSets counts the ways to draw exactly k non-overlapping line segments
// on n points (x = 0..n-1) such that each segment covers two or more points.
// Segments may share endpoints and do not have to cover all n points.
// The result is returned modulo 1e9+7.
//
// DP: f[i][j] = number of ways to place j segments using only the first i
// points (0..i-1), with f[i][0] = 1. The rightmost segment either does not end
// at point i-1 (f[i-1][j]), or it ends at i-1, starts at some point a < i-1,
// and the remaining j-1 segments lie within the first a+1 points:
//
//	f[i][j] = f[i-1][j] + sum_{a=0}^{i-2} f[a+1][j-1]
//
// A prefix sum over a keeps the inner loop O(k), so the total is O(n*k) time
// and O(k) space. The answer is f[n][k].
func numberOfSets(n int, k int) int {
	const mod = 1_000_000_007

	prev := make([]int, k+1)   // f[i-1][j]
	cur := make([]int, k+1)    // f[i][j]
	prefix := make([]int, k+1) // prefix[j] = sum_{a=1}^{i-1} f[a][j]

	prev[0] = 1 // f[0][0] = 1 (zero points, zero segments)

	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			v := prev[j]
			if j > 0 && i > 1 {
				v += prefix[j-1]
			}
			cur[j] = v % mod
		}
		for j := 0; j <= k; j++ {
			prefix[j] = (prefix[j] + cur[j]) % mod
		}
		prev, cur = cur, prev
	}

	return prev[k]
}
