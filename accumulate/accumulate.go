package accumulate

//Accumulate modifies given collection of strings using given converter.
func Accumulate(collection []string, converter func(string) string) []string {

	for i := range collection {
		collection[i] = converter(collection[i])
	}
	return collection

}
