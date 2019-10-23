package cryptosquare

import (
	"fmt"
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

type Rectangle [][]rune

func Encode(input string) string {
	input = normalize(input)

	c, r := 1, 0
	for {

		if c >= r && c-r <= 1 && c*r >= len(input) {
			break
		}
		c++
		r++
	}
	fmt.Println(c, r)
	return ""
}
