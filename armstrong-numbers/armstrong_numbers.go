package armstrong

import "math"

func IsNumber(number int) bool {
	digits := make([]float64, 0)
	numberOrginal := float64(number)
	var sum float64
	for number > 0 {
		digits = append(digits, float64(number%10))
		number /= 10

	}
	exponent := float64(len(digits))
	for _, digit := range digits {
		sum += math.Pow(digit, exponent)
	}
	return sum == numberOrginal
}
