package main

func countSeniors(details []string) int {
	count := 0
	for s := range details {
		if s[11] > '6'{
			count ++
		} else if s[11] == '6' && s[12] > '0'{
			count ++
		}
	}

	return count
}