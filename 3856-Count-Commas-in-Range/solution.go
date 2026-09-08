package leetcode

func countCommas(n int) int {
	total := 0
	lower, upper, digits := 1, 9, 1

	for lower <= n {
		hi := upper
		if hi > n {
			hi = n
		}

		count := hi - lower + 1
		commas := (digits - 1) / 3
		total += count * commas

		digits++
		lower = upper + 1
		upper = upper*10 + 9
	}

	return total
}
