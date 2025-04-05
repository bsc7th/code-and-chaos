// Install Go (if you haven't already)
// You can install it via HomeBrew "brew install go"


// To enable dependency tracking for your code by creating a go.mod file, run the go mod init command, giving it the name of the module your code will be in. The name is the module's module path.

// // This is your Go code. In this code, you:

// 1. Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
// 2. Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
// 3. Implement a main function to print a message to the console. A main function executes by default when you run the main package.

package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

// Run your code to see the greeting. "go run ."

// The go run command is one of many go commands you'll use to get things done with Go. Use the following command to get a list of the others:

// "go help"
