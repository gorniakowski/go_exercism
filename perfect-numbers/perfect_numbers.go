package perfect

import "errors"

type Classification int

var ErrOnlyPositive = errors.New("test")

const (
	ClassificationDeficient = iota
	ClassificationPerfect
	ClassificationAbundant
)

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}
	var aliquotSum int64 = 0

	for i := int64(1); int64(i) < n; i++ {
		if n%i == 0 {
			aliquotSum += i
		}
	}

	if aliquotSum == n {
		return ClassificationPerfect, nil
	}
	if aliquotSum > n {
		return ClassificationAbundant, nil
	}
	return ClassificationDeficient, nil

}
