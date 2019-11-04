package grains

import (
	"errors"
)

const chessBoardSize = 64

// Square gives a number of grains on a given chess field
func Square(n int) (uint64, error) {
	if n < 1 || n > chessBoardSize {
		return 0, errors.New("Value must be between 1 and 64")
	}

	return 1 << (n - 1), nil
}

//Total gives a total number of grains
func Total() uint64 {

	return 1<<(chessBoardSize) - 1
}
