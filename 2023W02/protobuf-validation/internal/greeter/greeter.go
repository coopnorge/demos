package greeter

import (
	"fmt"
)

// DoSomething does something.
func DoSomething(name string) (string, error) {
	if name == "John" {
		return "", fmt.Errorf("your name can't be %s", name)
	}
	return fmt.Sprintf("Hi %s!", name), nil
}
