package phonenumber

import "errors"

func Number(input string) (string, error) {
	result := make([]byte, 0)
	for _, char := range []byte(input) {
		if len(result) == 0 && char == 49 {
			continue
		}
		if char >= 48 && char <= 57 {
			result = append(result, char)
		}

		if len(result) >= 4 && (result[0] == 48 || result[3] < 50) {
			return "000", errors.New("wrong input")
		}
	}

	if len(result) != 10 {
		return "000", errors.New("wrong input,baby")
	}

	return string(result), nil
}

func AreaCode(input string) (string, error) {
	number, err := Number(input)
	return number[0:3], err
}
func Format(input string) (string, error) {
	number, err := Number(input)
	if err != nil {
		return "000", err
	}
	return "(" + number[0:3] + ")" + " " + number[3:6] + "-" + number[6:], nil
}
