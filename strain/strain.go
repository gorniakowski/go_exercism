package strain

type Ints []int
type Lists [][]int
type Strings []string

func (collction Ints) Keep(f func(int) bool) Ints {
	if collction == nil {
		return nil
	}
	result := make(Ints, 0)
	for _, item := range collction {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}
func (collction Ints) Discard(f func(int) bool) Ints {
	if collction == nil {
		return nil
	}
	result := Ints{}
	for _, item := range collction {
		if !f(item) {
			result = append(result, item)
		}
	}
	return result
}
func (collction Lists) Keep(f func([]int) bool) Lists {
	result := Lists{}
	for _, item := range collction {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}
func (collction Strings) Keep(f func(string) bool) Strings {
	result := Strings{}
	for _, item := range collction {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}
