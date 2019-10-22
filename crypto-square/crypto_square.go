package cryptosquare

import "strings"

func normalize(input string) (result string) {
	result = strings.ReplaceAll(strings.ToLower(input), " ", "")
	return
}
