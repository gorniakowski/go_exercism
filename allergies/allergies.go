package allergies

var allerNames = map[int]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

func Allergies(score uint) []string {
	scoreInt := int(score)
	result := []string{}
	var curr int = 0
	for scoreInt > 0 {

		for i := 0; i <= 7; i++ {

			curr = 1 << i
			if curr == scoreInt {
				break
			}
			if curr > scoreInt {
				curr = 1 << (i - 1)
				break
			}

		}
		scoreInt -= curr
		if scoreInt < 129 {
			result = append([]string{allerNames[curr]}, result...)
		}

	}
	return result
}

func AllergicTo(score uint, substance string) bool {
	alergies := Allergies(score)
	for _, allergy := range alergies {
		if allergy == substance {
			return true
		}
	}
	return false

}
