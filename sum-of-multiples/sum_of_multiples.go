package summultiples

//IsDivisible checks if divident is divvisible by any of the divisors
func IsDivisible(divident int, divisors []int) bool {
	for _, divisor := range divisors {
		if divisor != 0 && divident%divisor == 0 {
			return true
		}
	}
	return false
}

//SumMultiples return sum of the unique multiples
func SumMultiples(limit int, divisors ...int) (sum int) {

	for i := 1; i < limit; i++ {
		if IsDivisible(i, divisors) {
			sum += i
		}

	}
	return
}
