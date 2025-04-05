// Hello World in Go
// Go is a statically typed, compiled programming language designed by Google

package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

// Explanation:
// package main: In Go, the main package is the entry point of a Go program. It indicates that this file will contain the main function.
// import "fmt": We import the fmt package, which is used for formatted I/O operations, like printing to the console.
// func main() {}: The main function is the starting point of the Go program.
// fmt.Println("Hello, World!"): This prints "Hello, World!" to the console.

// **Variables and Types**

// Go is statically typed, which means you must declare the type of a variable when you create it.

package main

import "fmt"

func main() {
    var x int = 5
    var y string = "Hello, Go!"

    fmt.Println(x)
    fmt.Println(y)
}

// Explanation:
// var x int = 5: Here, x is a variable of type int (integer) and initialized to 5.
// var y string = "Hello, Go!": Similarly, y is a variable of type string with the value "Hello, Go!".
// fmt.Println(x): Prints the value of x (5).
// fmt.Println(y): Prints the value of y ("Hello, Go!").

// **Short Variable Declaration**

// Go also has a shorthand for declaring variables using :=:

package main

import "fmt"

func main() {
    x := 5      // shorthand for var x int = 5
    y := "Go!"  // shorthand for var y string = "Go!"

    fmt.Println(x)
    fmt.Println(y)
}

// x := 5 means "declare x as an integer with a value of 5".
// Go automatically infers the type from the right-hand side of the assignment.

// **Data Types in Go**

// Go supports several built-in types. Here are a few examples:

// int: Integer numbers (e.g., int, int32, int64)
// float64: Floating-point numbers (e.g., float32, float64)
// string: A sequence of characters (text)
// bool: Boolean values (true or false)

// **Control Structures**

// Just like many other languages, Go uses control structures like if-else and loops. Here's an example with if:

package main

import "fmt"

func main() {
    x := 10
    if x > 5 {
        fmt.Println("x is greater than 5")
    } else {
        fmt.Println("x is less than or equal to 5")
    }
}

// **Functions in Go**

// Go has a simple function definition syntax. Here’s a basic function that takes two integers and returns their sum:

package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(3, 4)
    fmt.Println(result)  // Output: 7
}

// func add(a int, b int) int defines a function called add that takes two arguments of type int and returns an int.
// return a + b returns the sum of the two input numbers.

// **Arrays and Slices**
// Arrays in Go have a fixed size, but slices are more flexible. Let's look at both:

package main

import "fmt"

func main() {
    var arr [3]int = [3]int{1, 2, 3}
    fmt.Println(arr)  // Output: [1 2 3]
}

// Slices:
// Slices are more commonly used because they can grow and shrink dynamically.

package main

import "fmt"

func main() {
    slice := []int{1, 2, 3}
    slice = append(slice, 4)  // Adds 4 to the slice
    fmt.Println(slice)         // Output: [1 2 3 4]
}

// append(slice, 4) adds 4 to the slice.
