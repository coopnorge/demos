package main

import "fmt"

// https://go.dev/doc/faq#nil_error

type Animal interface {
	Speak()
}

type Dog struct{}

func (d *Dog) Speak() {
	fmt.Printf("Woof!\n")
}

// GetAnimal returns an Animal interface.
// If returnTypedNil is true, it returns a nil *Dog pointer wrapped in the interface.
// If returnTypedNil is false, it returns a true nil interface.
func GetAnimal(returnTypedNil bool) Animal {
	if returnTypedNil {
		var myDog *Dog
		// TODO: A lot of lines that does something...
		// TODO: A lot of lines that does something...
		// TODO: A lot of lines that does something...
		// TODO: A lot of lines that does something...
		return myDog // This wraps the nil pointer in the interface
	}
	return nil
}

func main() {
	// Scenario 1: Returning a nil *Dog
	fmt.Println("Scenario 1: Returning a typed nil (*Dog)")
	a1 := GetAnimal(true)

	if a1 == nil {
		fmt.Println("a1 is nil")
		fmt.Printf("Value: %v, Type: %T\n", a1, a1)
	} else {
		fmt.Println("a1 is NOT nil (Trap!)")
		fmt.Printf("Value: %v, Type: %T\n", a1, a1)
		// a1.Speak()
	}

	fmt.Println("------------------------------------------------")

	// Scenario 2: Returning an explicit nil
	fmt.Println("Scenario 2: Returning explicit nil")
	a2 := GetAnimal(false)

	if a2 == nil {
		fmt.Println("a2 is nil (Correct)")
		fmt.Printf("Value: %v, Type: %T\n", a2, a2)
	} else {
		fmt.Println("a2 is NOT nil")
		fmt.Printf("Value: %v, Type: %T\n", a2, a2)
	}
}
