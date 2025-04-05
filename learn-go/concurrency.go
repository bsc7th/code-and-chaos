// Concurrency with Goroutines
// Go has built-in support for concurrency through goroutines. A goroutine is a lightweight thread of execution. You can run functions concurrently by using the go keyword.

// Here’s an example of launching a goroutine:

package main

import (
    "fmt"
    "time"
)

func sayHello() {
    time.Sleep(1 * time.Second)
    fmt.Println("Hello from goroutine!")
}

func main() {
    go sayHello()  // Launching sayHello as a goroutine
    
    fmt.Println("Main function is running concurrently.")
    time.Sleep(2 * time.Second)  // Give the goroutine time to complete
}

// Explanation:
// go sayHello() runs the sayHello function in a new goroutine concurrently with the rest of the program.
// time.Sleep(2 * time.Second) is used to ensure the main function waits long enough for the goroutine to complete.
