package anagram

import (
	"reflect"
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) (result []string) {

	for i := range candidates {
		if len(subject) != len(candidates[i]) {
			continue
		}
		if IsAnagram(subject, candidates[i]) {
			result = append(result, candidates[i])
		}
	}
	return
}

func IsAnagram(subject string, candidate string) bool {
	a := strings.Split(strings.ToLower(subject), "")
	b := strings.Split(strings.ToLower(candidate), "")

	if reflect.DeepEqual(a, b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)

	if reflect.DeepEqual(a, b) {
		return true
	}
	return false

}
