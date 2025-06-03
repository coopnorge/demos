package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var err error
	_ = err

	err = someFunction()
	pe := &CustomError{}
	if errors.As(err, &pe) {
		log.Printf("Got CustomError with value: %s", pe.Value)
	} else {
		panic("Expected errors.As to match the custom error")
	}

	log.Printf("Done.")
}

func someFunction() error {
	return &CustomError{
		Value: "some value",
	}
}

// CustomError implements the error interface as a receiver on a POINTER to the struct.
type CustomError struct {
	Value string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("some custom error with additional property: %s", e.Value)
}
