package main

import (
	"errors"
	"log"
)

func main() {
	var err error
	_ = err

	err = valueErrAsValue()
	if !errors.As(err, &ValueError{}) {
		panic("valueErrAsValue was not matched with errors.As (second param is pointer)")
	}
	if !errors.As(err, new(*ValueError)) {
		// panic("valueErrAsValue was not matched with errors.As (second param is double pointer)") // fails
	}

	err = valueErrAsPointer()
	if !errors.As(err, &ValueError{}) {
		// panic("valueErrAsPointer was not matched with errors.As (second param is pointer)") // fails
	}
	if !errors.As(err, new(*ValueError)) {
		panic("valueErrAsPointer was not matched with errors.As (second param is double pointer)")
	}

	valueErr := &ValueError{}
	if errors.As(err, &valueErr) {
		log.Printf("Got ValueError with value: %s", valueErr.Value)
	}

	log.Printf("Done.")
}

func valueErrAsPointer() error {
	return &ValueError{"foobar"}
}

func valueErrAsValue() error {
	return ValueError{"foobar"}
}
