package leetcode

func kthDistinct(arr []string, k int) string {
	m := make(map[string]int, len(arr))

	for _, s := range arr {
		m[s]++
	}
	index := 1

	for _, s := range arr {
		v := m[s]
		if v == 1 {
			if index == k {
				return s
			}
			index++
		}
	}

	return ""
}
