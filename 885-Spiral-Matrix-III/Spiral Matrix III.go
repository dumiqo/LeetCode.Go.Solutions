package leetcode

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	count := rows * cols
	result := [][]int{}
	result = append(result, []int{rStart, cStart})
	currentR := rStart
	currentC := cStart
	steps := 1
	increment := true

	for len(result) < count {
		if increment {
			for i := 0; i < steps; i++ {
				currentC++
				if len(result) >= count {
					break
				}

				if currentC < cols && currentC >= 0 && currentR < rows && currentR >= 0 {
					result = append(result, []int{currentR, currentC})
				}
			}
			for i := 0; i < steps; i++ {
				currentR++
				if len(result) >= count {
					break
				}

				if currentC < cols && currentC >= 0 && currentR < rows && currentR >= 0 {
					result = append(result, []int{currentR, currentC})
				}
			}
		} else {
			for i := 0; i < steps; i++ {
				currentC--
				if len(result) >= count {
					break
				}

				if currentC < cols && currentC >= 0 && currentR < rows && currentR >= 0 {
					result = append(result, []int{currentR, currentC})
				}
			}
			for i := 0; i < steps; i++ {
				currentR--
				if len(result) >= count {
					break
				}

				if currentC < cols && currentC >= 0 && currentR < rows && currentR >= 0 {
					result = append(result, []int{currentR, currentC})
				}
			}
		}
		increment = !increment
		steps++
	}

	return result
}
