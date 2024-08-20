package leetcode

func stoneGame(n []int) bool {
	alice := 0
	bob := 0
	if len(n)%2 != 0 {
		mid := len(n) / 2
		bob += n[mid]
		n = remove(n, mid)
	}

	for len(n) > 0 {
		mid := (len(n) / 2) - 1
		alice += maxOf(n[mid], n[mid+1])
		bob += minOf(n[mid], n[mid+1])

		n = remove(n, mid)
		n = remove(n, mid)
	}

	return alice > bob
}
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
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
