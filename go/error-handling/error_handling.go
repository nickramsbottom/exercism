package erratum

// Use an input with a resource from the opener
func Use(o ResourceOpener, input string) (result error) {
	resource, err := o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		resource, err = o()
	}

	defer resource.Close()

	defer func() {
		if r := recover(); r != nil {
			if frob, ok := r.(FrobError); ok {
				resource.Defrob(frob.defrobTag)
			}
			result = r.(error)
		}
	}()

	resource.Frob(input)
	return nil
}
