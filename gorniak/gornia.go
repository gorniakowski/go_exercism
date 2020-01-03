package main

import "fmt"

func main() {
	fmt.Println(strStr("a", "a"))
}

func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	if needleLen == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-needleLen; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}
