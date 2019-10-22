package pythagorean

type Triplet [3]int

func Range(min, max int) []Triplet {
	result := make([]Triplet, 0)
	for a := min; a <= max; a++ {
		for b := min; b <= max; b++ {
			for c := min; c <= max; c++ {
				if (a*a+b*b == c*c) && a < b && b < c {
					result = append(result, Triplet{a, b, c})
				}
			}
		}
	}
	return result
}

func Sum(p int) []Triplet {
	min := 1
	max := p / 2
	result := make([]Triplet, 0)
	triplets := Range(min, max)
	for _, triplet := range triplets {

		if triplet[0]+triplet[1]+triplet[2] == p {
			result = append(result, triplet)
		}
	}
	return result
}
