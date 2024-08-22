package leetcode

import (
	"strconv"
)

func findComplement(num int) int {
	n := int64(num)
	str := strconv.FormatInt(n, 2)
	complementStr := ""

	for _, c := range str {
		if c == '1' {
			complementStr += "0"
		} else {
			complementStr += "1"
		}
	}
	i, _ := strconv.ParseInt(complementStr, 2, 64)

	return int(i)
}
