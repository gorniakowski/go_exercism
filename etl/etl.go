package etl

import "strings"

type given map[int][]string
type expectation map[string]int

func Transform(input given) expectation {

	var output expectation = make(map[string]int)

	for score, lettersOld := range input {
		for _, letter := range lettersOld {
			output[strings.ToLower(letter)] = score
		}
	}

	return output
}
