// Pointers
// In Go, variables are passed by value by default, but you can also work with pointers to refer to the memory address of a value.

// Here’s an example of working with pointers:

package main

import "fmt"

func main() {
    x := 5
    p := &x  // Get the pointer to the memory address of 'x'
    
    fmt.Println("Value of x:", x)  // Output: 5
    fmt.Println("Memory address of x:", p)  // Output: <memory address>
    fmt.Println("Value at memory address p:", *p)  // Output: 5
}

// Explanation:

// p := &x gets the pointer to the memory address of x.
// *p dereferences the pointer, meaning it retrieves the value stored at the memory address.
