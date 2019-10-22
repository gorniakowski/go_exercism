package hamming

import (
	"errors"
)

//Distance calcutes number of different characters in two strigns
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("inputs have different lenght")
	}
	var counter = 0

	var aconv = []rune(a) // we convert string into runes to suport unicode
	var bconv = []rune(b)

	for i := range aconv {
		if aconv[i] != bconv[i] {
			counter++
		}

	}

	return counter, nil

}
