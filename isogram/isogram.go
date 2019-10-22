package isogram

import (
	"strings"
	"unicode"
)

//IsIsogram checks if the string is an isogram
func IsIsogram(input string) bool {

	input = strings.ToLower(input)
	letters := make(map[rune]bool)
	for _, letter := range input {

		if !unicode.IsLetter(letter) {
			continue
		}
		if letters[letter] {
			return false
		}
		letters[letter] = true

	}

	return true
}
