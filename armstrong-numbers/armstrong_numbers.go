package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(number int) bool {
	numberString := strconv.Itoa(number)
	var sum int
	for _, char := range numberString {
		digit, _ := strconv.Atoi(string(char))
		sum += int(math.Pow(float64(digit), float64(len(numberString))))
	}
	return sum == number
}
