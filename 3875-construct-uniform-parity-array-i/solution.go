package leetcode

func uniformArray(nums1 []int) bool {
	evenChan := make(chan bool)
	oddChan := make(chan bool)

	go isEven(nums1, evenChan)
	go isOdd(nums1, oddChan)

	even := <-evenChan
	odd := <-oddChan

	return even || odd
}

func isEven(nums1 []int, done chan<- bool) {
	for i := range nums1 {
		if nums1[i]%2 == 0 {
			continue
		}
		haveEven := false
		for j := range nums1 {
			if i == j {
				continue
			}

			if (nums1[i]-nums1[j])%2 == 0 {
				haveEven = true
				break
			}
		}
		if !haveEven {
			done <- false
		}
	}
	done <- true
}

func isOdd(nums1 []int, done chan<- bool) {
	for i := range nums1 {
		if nums1[i]%2 != 0 {
			continue
		}
		haveOdd := false
		for j := range nums1 {
			if i == j {
				continue
			}

			if (nums1[i]-nums1[j])%2 != 0 {
				haveOdd = true
				break
			}
		}
		if !haveOdd {
			done <- false
		}
	}
	done <- true

}
