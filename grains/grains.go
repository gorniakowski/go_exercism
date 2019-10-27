package grains

import "errors"

const chessBoardSize = 64

// Square gives a number of grains on a given chess field
func Square(n int) (uint64, error) {
	if n < 1 || n > chessBoardSize {
		return 0, errors.New("Value must be between 1 and 64")
	}
	var grains uint64 = 1

	for i := 1; i < n; i++ {
		grains *= 2
	}
	return grains, nil
}

//Total gives a total number of grains
func Total() uint64 {

	var sum uint64 = 0
	NoErrorSquare := func(n int) uint64 {
		result, _ := Square(n)
		return result
	}
	for i := 1; i <= chessBoardSize; i++ {
		sum += NoErrorSquare(i)
	}
	return sum
}
