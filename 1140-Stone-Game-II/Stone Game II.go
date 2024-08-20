package leetcode

func stoneGameII(n []int) int {
	l := len(n)
	sum := make([]int, l)
	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, l)
	}

	sum[l-1] = n[l-1]

	for i := l - 2; i >= 0; i-- {
		sum[i] = sum[i+1] + n[i]
	}

	return dfs(n, dp, 0, 1, sum)
}

func dfs(n []int, dp [][]int, i int, m int, sum []int) int {
	if i == len(n) {
		return 0
	}
	if i+2*m > len(n) {
		return sum[i]
	}
	if dp[i][m] != 0 {
		return dp[i][m]
	}

	min := sum[0] + 1000
	for x := 1; x <= 2*m; x++ {
		min = minOf(min, dfs(n, dp, i+x, maxOf(x, m), sum))
	}
	res := sum[i] - min
	dp[i][m] = res
	return res
}

func minOf(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func maxOf(x, y int) int {
	if x < y {
		return y
	}
	return x
}
