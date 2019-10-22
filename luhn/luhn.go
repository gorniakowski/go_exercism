package luhn

import (
	"strings"
)

//DoubleDigit doubles the given number if doubled number is over 9 it substracts 9 from it
func DoubleDigit(digit int) int {
	digit = digit * 2
	if digit > 9 {
		digit = digit - 9
	}
	return digit
}

//Valid return wherer the given number is valid per Luhns formula.
func Valid(input string) bool {

	input = strings.ReplaceAll(input, " ", "")
	if len(input) <= 1 {
		return false
	}
	sum := 0
	shallDouble := false
	for i := len(input) - 1; i >= 0; i-- {
		digit := int(input[i] - '0')

		if digit >= 10 {
			return false
		}
		if shallDouble {
			sum += DoubleDigit(digit)
			shallDouble = !shallDouble
			continue
		}
		sum += digit
		shallDouble = !shallDouble

	}

	return sum%10 == 0
}
