package erratum

import (
	"fmt"
)

type myResource struct {
	frobInput string
}

func (r myResource) Close() error {
	e := r.Close()
	if e != nil {
		return e
	}
	return nil
}

func (r myResource) Frob(input string) {
	r.frobInput = input
}

func (r myResource) Defrob(tag string) {
	r.frobInput = tag
}

func Use(o ResourceOpener, input string) error {
	res, err := o()

	// function that recovers panic
	defer res.Close()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Call recover in defer function, r = ", r)
			switch v := r.(type) {
			case FrobError:
				fmt.Println("FrobError")
				res.Defrob(v.defrobTag)
			case TransientError:
				fmt.Println("TransientError")
			}
		}

	}()

	fmt.Println("input", input)
	if err == nil {
		fmt.Println("1")
		res.Frob(input)
		// if err1 != nil {
		// 	frobErr, isFrobError := err1.(FrobError)
		// 	if isFrobError == true {
		// 		res.Defrob(frobErr.defrobTag)
		// 	}
		// 	return err1
		// }
	} else {
		switch v := err.(type) {
		case TransientError:
			fmt.Println(v)
			return err
		default:
			return err
		}
	}

	return nil
}
