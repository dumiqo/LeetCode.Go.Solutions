package leetcode

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestStableIndexExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"example 1", []int{5, 0, 1, 4}, 3, 3},
		{"example 2", []int{3, 2, 1}, 1, -1},
		{"example 3 (single element)", []int{0}, 0, 0},
		{"single element always stable", []int{7}, 0, 0},
		{"single element generous k", []int{7}, 100, 0},
		{
			"array longer than k is always stable at 0", []int{1, 2, 3, 4, 5}, 0, 0,
		},
		{"strictly decreasing never stable", []int{5, 4, 3, 2, 1}, 0, -1},
		{
			"strictly decreasing with large k", []int{5, 4, 3, 2, 1}, 4, 0,
		},
		{"all equal", []int{1, 1, 1, 1}, 0, 0},
		{"k large enough for first", []int{9, 0}, 9, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, smallestStableIndex(tt.nums, tt.k))
		})
	}
}

// TestSmallestStableIndexAgainstBruteForce cross-checks the O(n) approach
// against an independent oracle that recomputes max(nums[0..i]) and
// min(nums[i..n-1]) from scratch for every index i.
func TestSmallestStableIndexAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	for iter := 0; iter < 200; iter++ {
		n := rng.Intn(12) + 1
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rng.Intn(21) // keep within stated constraints (>= 0)
		}
		k := rng.Intn(11)

		want := bruteSmallestStableIndex(nums, k)
		got := smallestStableIndex(nums, k)
		assert.Equal(t, want, got, "nums=%v k=%d", nums, k)
	}
}

func bruteSmallestStableIndex(nums []int, k int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		prefMax := nums[0]
		for j := 1; j <= i; j++ {
			if nums[j] > prefMax {
				prefMax = nums[j]
			}
		}
		sufMin := nums[i]
		for j := i + 1; j < n; j++ {
			if nums[j] < sufMin {
				sufMin = nums[j]
			}
		}
		if prefMax-sufMin <= k {
			return i
		}
	}
	return -1
}
