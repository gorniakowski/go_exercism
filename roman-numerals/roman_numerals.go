package romannumerals

import "errors"

var translationTable = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func ToRomanNumeral(input int) (string, error) {

	if !(input > 0 && input < 3001) {
		return "", errors.New("Number out of range")
	}

	var output = ""
	for input > 0 {
		var max = 1
		for key := range translationTable {
			if key <= input && key > max {
				max = key
			}
		}
		output = output + translationTable[max]
		input = input - max

	}

	return output, nil
}
