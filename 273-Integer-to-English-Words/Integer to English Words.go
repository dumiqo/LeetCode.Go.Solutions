package leetcode

import (
	"fmt"
	"strings"
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	multiplier := 0
	del := map[int]string{
		0: "",
		1: "Thousand",
		2: "Million",
		3: "Billion",
	}

	result := ""
	for num > 0 {
		number := num % 1000
		numberStr := parse(number)
		if len(numberStr) > 0 {
			str := fmt.Sprintf("%s %s", numberStr, del[multiplier])
			str = strings.TrimSpace(str)
			result = fmt.Sprintf("%s %s", str, result)
			result = strings.TrimSpace(result)
		}
		num = num / 1000
		multiplier++
	}

	return result
}

func parse(num int) string {
	var result string

	if num >= 100 {
		result = getHundreds(num)
		num = num % 100
	}
	if num >= 20 {
		tens := getTens(num / 10)
		number := getNumber(num % 10)

		result = fmt.Sprintf("%s %s %s", result, tens, number)
		return strings.TrimSpace(result)
	}
	result = fmt.Sprintf("%s %s", result, getNumber(num))
	return strings.TrimSpace(result)
}

func getHundreds(num int) string {
	number := num / 100
	if number <= 0 {
		return ""
	}

	return getNumber(number) + " Hundred"
}

func getNumber(num int) string {
	switch num {
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	case 10:
		return "Ten"
	case 11:
		return "Eleven"
	case 12:
		return "Twelve"
	case 13:
		return "Thirteen"
	case 14:
		return "Fourteen"
	case 15:
		return "Fifteen"
	case 16:
		return "Sixteen"
	case 17:
		return "Seventeen"
	case 18:
		return "Eighteen"
	case 19:
		return "Nineteen"
	default:
		return ""
	}
}

func getTens(num int) string {
	switch num {
	case 2:
		return "Twenty"
	case 3:
		return "Thirty"
	case 4:
		return "Forty"
	case 5:
		return "Fifty"
	case 6:
		return "Sixty"
	case 7:
		return "Seventy"
	case 8:
		return "Eighty"
	case 9:
		return "Ninety"
	default:
		return ""
	}
}
