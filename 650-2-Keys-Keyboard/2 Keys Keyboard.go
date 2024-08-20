package leetcode

func minSteps(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = n + 1000
	}
	dp[1] = 0

	for i := 2; i <= n; i++ {
		dp[i] = min(i, dp[i])

		index := 2
		for j := i + i; j <= n; j += i {
			dp[j] = min(dp[i]+index, dp[j])
			index++
		}
	}

	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
