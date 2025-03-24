package main

import "fmt"

// PointerError implements the error interface as a receiver on a POINTER to the struct.
type PointerError struct {
	Value string
}

func (e *PointerError) Error() string {
	return fmt.Sprintf("some custom error with additional property: %s", e.Value)
}
