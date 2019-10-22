package wordcount

import (
	"strings"
	"unicode"
)

//Frequency represents frequency of words in  a string
type Frequency map[string]int

//GetWord extracs words from string retruring a slice of words
func GetWord(s string) []string {
	return strings.FieldsFunc(s, func(char rune) bool {
		return !(unicode.IsLetter(char) || unicode.IsDigit(char) || char == '\'')
	})
}

//TrimQuotations chcecs if word is in quatation  '' and removes them
func TrimQuotations(s string) string {
	if s[0] == '\'' && s[len(s)-1] == '\'' {
		return s[1 : len(s)-1]
	}
	return s
}

//WordCount counts words occrences in the text
func WordCount(input string) Frequency {
	result := make(Frequency)
	words := GetWord(input)
	for _, word := range words {
		word = strings.ToLower(TrimQuotations(word))
		result[word]++
	}
	return result
}
