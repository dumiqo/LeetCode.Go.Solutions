package leetcode

func sumGame(num string) bool {

	topSum, topQ := stat(num[:len(num)/2])
	bottomSum, bottomQ := stat(num[len(num)/2:])

	return false
}

func stat(num string) (int, int) {

	sum := 0
	qCount := 0
	for _, ch := range num {
		if ch == '?' {
			qCount++
		} else {
			var in = int(ch - '0')
			sum += in
		}
	}

	return sum, qCount
}
