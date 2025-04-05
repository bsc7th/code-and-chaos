// Interfaces
// An interface in Go is a type that specifies a set of methods. A type is said to implement an interface if it provides definitions for all the methods declared by that interface. Unlike many other languages, Go does not require explicit declarations to indicate that a type implements an interface; it is implicit.

// Here’s an example of defining and using an interface:

package main

import "fmt"

// Define an interface called 'Speaker'
type Speaker interface {
    Speak() string
}

// Define a struct 'Person'
type Person struct {
    Name string
}

// Method 'Speak' implements the 'Speaker' interface for 'Person'
func (p Person) Speak() string {
    return "Hello, my name is " + p.Name
}

func main() {
    person := Person{Name: "Charlie"}
    
    var speaker Speaker = person  // person satisfies the Speaker interface
    fmt.Println(speaker.Speak())  // Output: Hello, my name is Charlie
}

// Explanation:

// Speaker is an interface with a single method Speak() string.
// Person implements the Speak method, and thus satisfies the Speaker interface.
// We can assign a Person instance to a variable of type Speaker, and call the Speak method.
