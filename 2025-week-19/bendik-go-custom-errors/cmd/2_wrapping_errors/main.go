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
	if errors.Is(err, errDBUserNotFound) {
		log.Printf("Found specific errDBUserNotFound error, 3 layers deep!")
	} else {
		panic("Expected errors.Is to find the specific error")
	}

	log.Printf("Done.")
}

func someFunction() error {
	err := someDatabaseFunction()
	if err != nil {
		return fmt.Errorf("something went wrong when calling the database: %w", err)
	}
	return nil
}

func someDatabaseFunction() error {
	return fmt.Errorf("error when calling the database: %w", errDBUserNotFound)
}
