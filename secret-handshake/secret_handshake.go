package secret

// Handshake decodes the given code into the secret handshake steps
func Handshake(code uint) []string {
	handshake := []string{}

	if code&1 == 1 {
		handshake = append(handshake, "wink")
	}
	if code&2 == 2 {
		handshake = append(handshake, "double blink")
	}
	if code&4 == 4 {
		handshake = append(handshake, "close your eyes")
	}
	if code&8 == 8 {
		handshake = append(handshake, "jump")
	}
	if code&16 == 16 {
		for i := len(handshake)/2 - 1; i >= 0; i-- {
			opp := len(handshake) - 1 - i
			handshake[i], handshake[opp] = handshake[opp], handshake[i]
		}
	}

	return handshake
}
