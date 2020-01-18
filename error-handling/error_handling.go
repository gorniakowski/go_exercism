package erratum

import "fmt"

import "errors"

func Use(o ResourceOpener, input string) error {
	resource, err := o()
	var transErr TransientError
	if errors.As(err, &transErr) {
		Use(o, input)
	}
	if err != nil {
		return err
	}

	fmt.Println(resource, err)

	defer resource.Close()
	return nil
}
