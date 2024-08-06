package leetcode

import (
	"sort"
)

func minimumPushes(word string) int {
	chars := make(map[rune]int)

	for _, c := range word {
		chars[c]++
	}
	if len(chars) <= 8 {
		return len(word)
	}
	nums := make([]int, len(chars))

	for _, n := range chars {
		nums = append(nums, n)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	result := 0
	multiplier := 1
	count := 8
	for _, i := range nums {
		result += i * multiplier
		count--
		if count <= 0 {
			count = 8
			multiplier++
		}
	}

	return result
}
