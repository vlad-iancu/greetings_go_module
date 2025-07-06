package others

import (
	"fmt"
)

func OtherGreeting(name string) (string, error) {
	// If the name is empty, return an error.
	if name == "" {
		return "", fmt.Errorf("empty name")
	}

	// Create a message using the name.
	message := fmt.Sprintf("Hello, %v. Welcome!", name)

	// Return the message and a nil error.
	return message, nil
}