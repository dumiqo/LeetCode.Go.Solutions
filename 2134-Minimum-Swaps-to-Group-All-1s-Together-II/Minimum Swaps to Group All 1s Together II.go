package leetcode

func minSwaps(nums []int) int {
	n := len(nums)

	totalOnes := 0

	for _, i := range nums {
		if i == 1 {
			totalOnes++
		}
	}

	if totalOnes <= 1 {
		return 0
	}

	var prefixSum []int
	prefixSum = make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = nums[i] + prefixSum[i]
	}

	minSwaps := n + 3

	for i := 0; i < n; i++ {
		end := i + totalOnes - 1
		var currentOnes int

		if end < n {
			currentOnes = prefixSum[end+1] - prefixSum[i]
		} else {
			currentOnes = (prefixSum[n] - prefixSum[i]) + prefixSum[end-n+1]
		}
		newValue := totalOnes - currentOnes
		if minSwaps > newValue {
			minSwaps = newValue
		}
	}

	return minSwaps
}
