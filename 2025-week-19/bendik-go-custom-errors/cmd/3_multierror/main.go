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
	return errors.Join(err1, err2)
}

func someDatabaseFunction() error {
	return fmt.Errorf("error when calling the database: %w", errDBUserNotFound)
}

func someHTTPClientFunction() error {
	return fmt.Errorf("error when calling an HTTP endpoint")
}
