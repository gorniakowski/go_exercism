package sieve

import (
	"math"
)

type pair struct {
	number int
	mark   bool
}

func Sieve(limit int) (result []int) {

	if limit == 1 {
		return []int{}
	}
	table := []pair{}

	for i := 2; i <= limit; i++ {
		table = append(table, pair{i, true})

	}

	for i := 2; float64(i) <= math.Sqrt(float64(limit)); i++ {
		if table[i].mark {
			s := 0
			for j := i * i; j <= limit; j = (i * i) + (s * i) {
				s++
				table[j-2].mark = false
			}
		}
	}
	for _, item := range table {
		if item.mark {
			result = append(result, item.number)
		}
	}
	return
}
