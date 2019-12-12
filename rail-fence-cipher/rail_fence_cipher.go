package railfence

func Encode(message string, rails int) string {

	rail := make([]string, rails)
	var result string
	var railCounter int = 0
	var upOrDown = false
	for i := range message {
		rail[railCounter] = rail[railCounter] + string(message[i])
		if railCounter == rails-1 || railCounter == 0 {
			upOrDown = !upOrDown
		}
		if upOrDown {
			railCounter++
			continue
		}
		railCounter--

	}
	for i := range rail {
		result += rail[i]
	}
	return result
}

func Decode(message string, rails int) string {

	railLenght := make([]int, rails)
	fence := make([]string, rails)
	railCounter := 0
	upOrDown := false
	//calculate how long each has to be
	for range message {
		railLenght[railCounter]++
		if railCounter == rails-1 || railCounter == 0 {
			upOrDown = !upOrDown
		}
		if upOrDown {
			railCounter++
			continue
		}
		railCounter--

	}

	//populate all the rails
	railCounter = 0
	charCounter := 0
	for _, char := range message {
		fence[railCounter] += string(char)
		if charCounter == railLenght[railCounter]-1 {
			railCounter++
			charCounter = 0
			continue
		}
		charCounter++
	}
	//go along the zig-zag and get the message
	railCounter = 0
	upOrDown = false
	result := ""
	for range message {
		result += string(fence[railCounter][0])
		fence[railCounter] = fence[railCounter][1:]
		if railCounter == rails-1 || railCounter == 0 {
			upOrDown = !upOrDown
		}
		if upOrDown {
			railCounter++
			continue
		}
		railCounter--

	}

	return result

}
