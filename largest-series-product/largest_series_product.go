package lsproduct

import (
	"errors"
	"strconv"
	"strings"
)

func LargestSeriesProduct(series string, n int) (result int, err error) {

	if len(series) < n {
		return 0, errors.New("series must be longer than subseries")

	}
	if n < 0 {
		return 0, errors.New("Span can't be negative")
	}
	if n == 0 {
		return 1, nil
	}
	var numbersStr []string = strings.Split(series, "")
	var numbersInt = []int{}

	for _, numberStr := range numbersStr {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			return 0, errors.New("wrong series")
		}
		numbersInt = append(numbersInt, number)

	}
	for i := range numbersInt {
		product := 1
		if i+n > len(numbersInt) {
			break
		}
		for j := 0; j < n; j++ {
			product *= numbersInt[i+j]
		}
		if result < product {
			result = product
		}
	}

	return result, nil
}
