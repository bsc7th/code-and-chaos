package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name srting) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprinttf("Hi, %v. Welcome!", name)
	return message
}

// The code above might have been written this way
var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)
