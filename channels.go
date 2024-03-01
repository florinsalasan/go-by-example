package main

import "fmt"

// Channels are the pipes that can connect concurrent goroutines. You
// can send values into channels from one goroutine and receive those
// values into another goroutine.
func main() {

    // a channel is created using make(chan val-type), channels are typed
    // by the values they convey
    messages := make(chan string)

    // A value is sent into a channel using the channel <- syntax. Here we 
    // send "ping" to the messages channel made above from a new goroutine
    go func() { messages <- "ping" }()

    // the syntax <-channel receives a value from the channel, this is where
    // the "ping" message is received from the line that sent it to messages
    // above
    msg := <-messages
    fmt.Println(msg)

    // When running this the "ping" is successfully passed from one goroutine
    // to another via the channel

    // By default sends and receives block until both the sender and receiver
    // are ready. This property allowed us to wait at the end of the program
    // without having to use synchronization.

}
