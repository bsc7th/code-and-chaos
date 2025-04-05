// Channels
// Channels are used to communicate between goroutines in Go. A channel allows you to send and receive values between different goroutines.

// Here’s an example of using a channel:

package main

import "fmt"

func greet(ch chan string) {
    ch <- "Hello from goroutine!"  // Send message to channel
}

func main() {
    ch := make(chan string)  // Create a new channel of type string
    
    go greet(ch)  // Launch goroutine
    
    message := <-ch  // Receive message from the channel
    fmt.Println(message)  // Output: Hello from goroutine!
}

// Explanation:

// ch <- "Hello from goroutine!" sends a message into the channel.
// <-ch receives a message from the channel.
