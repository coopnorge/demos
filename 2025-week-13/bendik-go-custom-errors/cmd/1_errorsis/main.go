package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	errErrorsNew = errors.New("some error from errors.New")
	errFmtErrorf = fmt.Errorf("some error from fmt.Errorf")
)

func main() {
	var err error

	err = doSomethingErrorsNew()
	if !errors.Is(err, errErrorsNew) {
		panic("expected errors.Is")
	}
	if err != errErrorsNew { // Don't do this
		panic("expected equal errors")
	}

	err = doSomethingFmtErrorf()
	if !errors.Is(err, errFmtErrorf) {
		panic("expected errors.Is")
	}
	if err != errFmtErrorf { // Don't do this
		panic("expected equal errors")
	}

	err = withWrap(doSomethingErrorsNew)
	if !errors.Is(err, errErrorsNew) {
		panic("expected errors.Is")
	}
	if err != errErrorsNew { // Don't do this
		// panic("expected equal errors")
	}

	err = withWrap(doSomethingFmtErrorf)
	if !errors.Is(err, errFmtErrorf) {
		panic("expected errors.Is")
	}
	if err != errFmtErrorf { // Don't do this
		// panic("expected equal errors")
	}

	log.Printf("Done.")
}

func doSomethingErrorsNew() error {
	return errErrorsNew
}

func doSomethingFmtErrorf() error {
	return errFmtErrorf
}

func withWrap(f func() error) error {
	err := f()
	if err != nil {
		return fmt.Errorf("wrapping err: %w", err)
	}
	return nil
}
