package erratum

// Use an input with a resource from the opener
func Use(o ResourceOpener, input string) (e error) {
	var err error
	var res Resource

	defer func() {
		if r := recover(); r != nil {
			frobError, isFrobError := r.(FrobError)

			if isFrobError {
				res.Defrob(frobError.defrobTag)
			}

			e = r.(error)
			res.Close()
		}
	}()

	res, err = o()

	for err != nil && err.Error() == "some error" {
		res, err = o()
	}

	if err != nil && err.Error() != "some error" {
		return err
	}

	res.Frob(input)
	res.Close()
	return err
}
