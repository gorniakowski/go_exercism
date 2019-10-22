package cryptosquare

import "unicode"

func normalize(input string) (result string) {

	for _, char := range input {
		if unicode.IsDigit(char) && unicode.IsLetter(char) {
			char = unicode.ToLower(char)
			result += string(char)
		}
	}
	return
}
