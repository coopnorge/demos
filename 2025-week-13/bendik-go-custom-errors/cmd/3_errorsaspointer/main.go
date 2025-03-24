package main

import (
	"errors"
	"log"
)

func main() {
	var err error
	_ = err

	// err = pointerErrAsValue() // impossible

	err = pointerErrAsPointer()

	// Lint error ("go vet" and "errorsas"): second argument to errors.As must be a non-nil pointer to either a type that implements error, or to any interface type
	// if !errors.As(err, &PointerError{}) {
	// 	panic("pointerErrAsPointer was not matched with errors.As (second param is pointer)")
	// }

	if !errors.As(err, new(*PointerError)) {
		panic("pointerErrAsPointer was not matched with errors.As (second param is double pointer)")
	}

	log.Printf("Done.")
}

//func pointerErrAsValue() error {
//	return PointerError{} // does not compile
//}

func pointerErrAsPointer() error {
	return &PointerError{}
}
