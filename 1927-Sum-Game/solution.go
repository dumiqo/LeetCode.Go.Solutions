package leetcode

import (
	"math"
)

func sumGame(num string) bool {
	sum1, q1 := stat(num[:len(num)/2])
	sum2, q2 := stat(num[len(num)/2:])

	totalQ := q1 + q2
	if totalQ%2 == 1 { // нечётное число вопросов – последний ход за Алисой
		return true
	}

	if q1 > q2 && sum1 > sum2 {
		return true
	}

	if q1 < q2 && sum1 < sum2 {
		return true
	}

	diff := int(math.Abs(float64(sum1 - sum2)))
	diffQ := int(math.Abs(float64(q1 - q2)))

	// Если вопросы распределены поровну
	if diffQ == 0 {
		return diff != 0
	}

	// Боб может выиграть только при точном равенстве
	return diff != 9*(diffQ/2)
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
