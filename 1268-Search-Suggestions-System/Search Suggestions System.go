package leetcode

import (
	"sort"
	"strings"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Sort(sort.StringSlice(products))
	result := make([][]string, len(searchWord))
	var strs []string
	for i := 1; i <= len(searchWord); i++ {
		strs = append(strs, searchWord[0:i])
	}

	for i := 0; i < len(searchWord); i++ {
		str := strs[i]
		for _, prod := range products {
			if len(prod) <= 0 {
				continue
			}
			if len(result[i]) >= 3 {
				break
			}
			if strings.HasPrefix(prod, str) {
				result[i] = append(result[i], prod)
			}
		}
	}

	return result
}
