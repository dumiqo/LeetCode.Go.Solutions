package leetcode

import "sort"

func shortestBeautifulSubstring(s string, k int) string {

	subs := make(map[int][]string, 0)
	minLen := len(s)
	for left := 0; left < len(s); left++ {
		sum := 0
		for right := left; right < len(s); right++ {
			sum += getNum(s, right)
			if sum == k {
				str := s[left : right+1]
				if len(str) < minLen {
					minLen = len(str)
				}
				subs[len(str)] = append(subs[len(str)], str)
			} else if sum > k {
				break
			}
		}
	}

	if len(subs) == 0 {
		return ""
	}
	strs := subs[minLen]

	sort.Slice(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	})
	return strs[0]
}
func getNum(s string, i int) int {
	return int(s[i] - '0')
}
