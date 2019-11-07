package lsproduct

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func LargestSeriesProduct(series string, n int) (result int, err error) {

	if len(series) < n {
		return 0, errors.New("series must be longer than subseries")

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
	product := 0
	for i := range numbersInt {
		if i+n > len(numbersInt) {
			fmt.Println("you are here")
			break
		}
		for j := 0; j < n; j++ {
			product *= numbersInt[i+j]
		}
	}
	fmt.Println(product)
	return 0, nil
}
