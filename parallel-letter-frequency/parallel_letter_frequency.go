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
	chMap := make(chan FreqMap)

	for _, text := range input {
		go func(chMap chan FreqMap, text string) {
			chMap <- Frequency(text)
		}(chMap, text)
	}

	for range input {
		select {
		case r := <-chMap:
			for i := range r {
				result[i] += r[i]
			}
		}
	}
	return result
}
