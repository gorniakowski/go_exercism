package palindrome

import (
	"errors"
	"strconv"
)

type Product struct {
	Product        int
	Factorizations [][2]int
}

func reverse(s int) int {
	rs := []rune(strconv.Itoa(s))
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	result, _ := strconv.Atoi(string(rs))

	return result
}

func Products(fmin, fmax int) (pmin, pmax Product, err error) {

	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax...")
	}

	var min, max int = fmax * fmax, 0

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			prod := i * j
			if prod == reverse(prod) {
				if min > prod {
					min = prod
				}
				if max < prod {
					max = prod
				}
			}
		}

	}
	factorsMin := make([][2]int, 0)
	factorsMax := make([][2]int, 0)
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			if i*j == min {
				factorsMin = append(factorsMin, [2]int{i, j})
			}
			if i*j == max {
				factorsMax = append(factorsMax, [2]int{i, j})
			}

		}
	}
	if len(factorsMin) == 0 || len(factorsMax) == 0 {
		return Product{}, Product{}, errors.New("no palindromes...")
	}
	return Product{min, factorsMin}, Product{max, factorsMax}, nil
}
