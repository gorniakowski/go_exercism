package strain

type Ints []int
type Lists [][]int
type Strings []string

func (collction Ints) Keep(f func(int) bool) (result Ints) {

	for _, item := range collction {
		if f(item) {
			result = append(result, item)
		}
	}
	return
}
func (collction Ints) Discard(f func(int) bool) Ints {
	return collction.Keep(func(i int) bool { return !f(i) })
}
func (collction Lists) Keep(f func([]int) bool) (result Lists) {

	for _, item := range collction {
		if f(item) {
			result = append(result, item)
		}
	}
	return
}
func (collction Strings) Keep(f func(string) bool) (result Strings) {

	for _, item := range collction {
		if f(item) {
			result = append(result, item)
		}
	}
	return
}
