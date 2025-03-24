package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var err error
	_ = err

	err = typedErrorNotNil()
	if !errors.As(err, new(*PointerError)) {
		panic("typedErrorNotNil was not matched with errors.As (second param is double pointer)")
	}

	err = typedErrorNil()
	if err != nil {
		panic("typedErrorNil returned a typed nil error, which is not nil") // fails
	}

	log.Printf("Done.")
}

func typedErrorNotNil() *PointerError {
	return &PointerError{
		Value: "some value",
	}
}

func typedErrorNil() *PointerError {
	return nil
}

// PointerError implements the error interface as a receiver on a POINTER to the struct.
type PointerError struct {
	Value string
}

func (e *PointerError) Error() string {
	return fmt.Sprintf("some custom error with additional property: %s", e.Value)
}
