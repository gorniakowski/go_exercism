package prime

func Factors(input int64) []int64 {
	result := make([]int64, 0)
	for input > 1 {

		for i := int64(2); i <= input; i++ {
			if input%i == 0 {
				result = append(result, i)
				input /= i
				break
			}
		}
	}
	return result
}
