package leetcode

func lemonadeChange(bills []int) bool {
	bank := make(map[int]int)
	bank[5] = 0
	bank[10] = 0
	bank[5] = 0

	for _, bil := range bills {
		if bil == 5 {
			bank[5]++
		} else if bil == 10 {
			if bank[5] >= 1 {
				bank[5]--
				bank[10]++
			} else {
				return false
			}
		} else if bil == 20 {
			if bank[10] >= 1 && bank[5] >= 1 {
				bank[10]--
				bank[5]--
				bank[20]++
			} else if bank[5] >= 3 {
				bank[5] -= 3
				bank[20]++
			} else {
				return false
			}
		}
	}

	return true
}
