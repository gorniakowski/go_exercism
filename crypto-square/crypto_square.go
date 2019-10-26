package cryptosquare

import (
	"strings"
	"unicode"
)

func normalize(input string) (result string) {

	for _, char := range input {
		if unicode.IsDigit(char) || unicode.IsLetter(char) {
			char = unicode.ToLower(char)
			result += string(char)

		}
	}
	return
}

func Encode(input string) string {
	input = normalize(input)

	c, r := 1, 0
	alternator := true
	for {
		if c >= r && c-r <= 1 && c*r >= len(input) {
			break
		}
		if alternator {
			r++
			alternator = !alternator
		} else {
			c++
			alternator = !alternator
		}

	}
	for {
		if len(input) < c*r {
			input = input + " "
		} else {
			break
		}
	}

	square := []string{}
	result := ""
	for i := 0; i < r; i++ {
		row := input[i*c : i*c+c]
		square = append(square, row)
	}

	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			result += string(square[j][i])
		}
		result += " "
	}

	return strings.TrimSuffix(result, " ")
}
