// **Structs**
// In Go, structs are used to group related data together.
// A struct is a composite data type that groups fields of different types. Think of a struct as similar to a class in object-oriented programming but without inheritance.

// example of struct

package main

import "fmt"

// Define a struct called 'Person'
type Person struct {
    Name string
    Age  int
}

func main() {
    // Create an instance of 'Person'
    person := Person{Name: "Alice", Age: 30}
    fmt.Println(person)  // Output: {Alice 30}
}

// Explanation:

// type Person struct {} defines a new struct called Person with two fields: Name (a string) and Age (an integer).
// person := Person{Name: "Alice", Age: 30} creates a new instance of Person, initializing Name to "Alice" and Age to 30.
// fmt.Println(person) prints the Person struct instance, displaying {Alice 30}.

// **Methods on Structs**
// In Go, you can define methods on structs. A method is a function that is associated with a type (like a struct). The method has a special receiver argument.

// Here’s how we add a method to the Person struct:

package main

import "fmt"

// Define a struct called 'Person'
type Person struct {
    Name string
    Age  int
}

// Method that belongs to the 'Person' struct
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}

func main() {
    person := Person{Name: "Bob", Age: 25}
    person.Greet()  // Output: Hello, my name is Bob
}

// Explanation:

// func (p Person) Greet() defines a method Greet on the Person type. p is the receiver of type Person, which is the instance the method operates on.
// person.Greet() calls the Greet method for the person instance.
