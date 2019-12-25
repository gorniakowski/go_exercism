package main

import "fmt"

func main() {
	fmt.Println(romanToInt("IV"))
}

var values = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var result int
	for i, sym := range s {
		curr := values[byte(sym)]
		if i <= len(s)-2 && curr < values[s[i+1]] {
			result -= curr
			continue
		}
		result += curr
	}
	return result
}
