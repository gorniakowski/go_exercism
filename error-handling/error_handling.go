package erratum

//Use opens Resourse and runs Frob() on in handling all errors gracefully
func Use(o ResourceOpener, input string) (result error) {

	resource, err := o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		resource, err = o()

	}
	defer func() {
		if r, ok := recover().(error); ok {
			if r, ok := r.(FrobError); ok {
				resource.Defrob(r.defrobTag)
			}
			result = r
		}
		resource.Close()
	}()
	resource.Frob(input)
	return result
}
