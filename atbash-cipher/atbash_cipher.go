package atbash

import "unicode"

import "strings"

func Atbash(input string) string {
	var result string
	var count int = 0
	for i, char := range input {

		if !(unicode.IsLetter(char) || unicode.IsDigit(char)) {
			continue
		}

		if char >= 65 && char <= 90 {
			char = unicode.ToLower(char)
		}
		if char >= 97 && char <= 122 {
			result += string(219 - char)
		} else {
			result += string(char)
		}

		count++
		if count > 4 && i < len(input)-1 {
			count = 0
			result += " "
		}

	}
	return strings.Trim(result, " ")
}
