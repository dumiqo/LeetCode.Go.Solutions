package leetcode

func uniformArray(nums1 []int) bool {
	allOdd, allEvent := true, true
	for i := range nums1 {
		haveEven := nums1[i]%2 == 0
		haveOdd := nums1[i]%2 != 0
		for j := range nums1 {
			if i == j {
				continue
			}
			if haveEven == haveOdd && haveOdd == true {
				break
			}

			if haveEven == false && (nums1[i]-nums1[j])%2 == 0 {
				haveEven = true
			}
			if haveOdd == false && (nums1[i]-nums1[j])%2 != 0 {
				haveOdd = true
			}
		}
		allOdd = haveOdd
		allEvent = haveEven
	}
	return allOdd || allEvent
}
