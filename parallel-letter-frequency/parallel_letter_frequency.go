package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

//ConcurrentFrequency returns character frequency map form slice of strings
func ConcurrentFrequency(input []string) FreqMap {
	result := make(FreqMap)
	chMap := make(chan FreqMap, 10)

	for _, text := range input {
		go func(text string) {
			chMap <- Frequency(text)
		}(text)
	}

	for range input {
		for k, v := range <-chMap {
			result[k] += v
		}
	}

	return result
}
