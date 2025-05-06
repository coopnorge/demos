package main

import (
	"errors"
	"fmt"
	"log"
)

var errDBUserNotFound error = errors.New("user not found in database")

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

	if errors.Is(err, errDBUserNotFound) {
		log.Printf("Found specific errDBUserNotFound error!")
	} else {
		panic("Expected errors.Is to find the specific error")
	}

	log.Printf("Done.")
}

func someFunction() error {
	// Note: Not returning early if there is an error
	err1 := someDatabaseFunction()
	err2 := someHTTPClientFunction()

	return &CustomError{
		Value:  "some value",
		Errors: []error{err1, err2},
	}
}

func someDatabaseFunction() error {
	return fmt.Errorf("error when calling the database: %w", errDBUserNotFound)
}

func someHTTPClientFunction() error {
	return fmt.Errorf("error when calling an HTTP endpoint")
}

// CustomError implements the error interface as a receiver on a POINTER to the struct.
type CustomError struct {
	Value  string
	Errors []error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("some custom error with additional property: %s", e.Value)
}

// "Unwrap() []error" and "Unwrap() error" are "hidden interfaces".
// https://cs.opensource.google/go/go/+/refs/tags/go1.24.2:src/errors/wrap.go;l=61-76
func (e *CustomError) Unwrap() []error {
	return e.Errors
}
