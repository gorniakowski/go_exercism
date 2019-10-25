package luhn

import (
	"strings"
)

//Valid return wherer the given number is valid per Luhns formula.
func Valid(input string) bool {

	input = strings.ReplaceAll(input, " ", "")
	if len(input) <= 1 {
		return false
	}
	sum := 0
	shallDouble := len(input)%2 == 0
	for _, digit := range input {
		digit := int(digit - '0')
		if digit >= 10 {
			return false
		}

		if shallDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		shallDouble = !shallDouble

	}

	return sum%10 == 0
}
