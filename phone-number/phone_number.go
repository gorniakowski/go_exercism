package phonenumber

func Number(input string) string {
	result := make([]byte, 0)
	for _, char := range []byte(input) {
		if char >= 48 && char <= 57 {
			result = append(result, char)
		}
	}
	return string(result)
}
