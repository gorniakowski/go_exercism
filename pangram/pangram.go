package pangram

import (
	"log"
	"regexp"
	"strings"
)

func IsPangram(input string) bool {

	var inputLetters = make(map[rune]int)
	reg, err := regexp.Compile("[^a-zA-Z]")
	if err != nil {
		log.Fatal(err)
	}
	input = strings.ToLower(reg.ReplaceAllLiteralString(input, ""))
	for _, letter := range input {
		inputLetters[letter]++

	}

	if len(inputLetters) == 26 {
		return true
	}
	return false
}
