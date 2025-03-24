package main

import "fmt"

// ValueError implements the error interface as a receiver on a VALUE to the struct.
type ValueError struct {
	Value string
}

func (e ValueError) Error() string {
	return fmt.Sprintf("some custom error with additional property: %s", e.Value)
}
