package leetcode

// smallestStableIndex returns the smallest index i such that the instability
// score max(nums[0..i]) - min(nums[i..n-1]) is less than or equal to k,
// or -1 if no such index exists.
//
// The prefix maximum and suffix minimum arrays are each built in one pass,
// and the scan happens in another, giving O(n) time and O(n) space.
func smallestStableIndex(nums []int, k int) int {
	n := len(nums)

	prefMax := make([]int, n)
	prefMax[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > prefMax[i-1] {
			prefMax[i] = nums[i]
		} else {
			prefMax[i] = prefMax[i-1]
		}
	}

	sufMin := make([]int, n)
	sufMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < sufMin[i+1] {
			sufMin[i] = nums[i]
		} else {
			sufMin[i] = sufMin[i+1]
		}
	}

	for i := 0; i < n; i++ {
		if prefMax[i]-sufMin[i] <= k {
			return i
		}
	}
	return -1
}
