package main

import "fmt"

// when using channels as function parameters, it can be specified if a 
// channel is meant to only send or receive values. This specificity 
// increases the type-safety of the written code


// this ping function only accepts a channel for sending values, it would
// be a compile-time error to try and receive on this channel
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// the pong function accepts one channel for receives (pings) and a sceond
// channel for sends (pongs)
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {

    pings := make(chan string, 1)
    pongs := make(chan string, 1)

    ping(pings, "passed message")
    // in the line above, msg is sent to pings
    pong(pings, pongs)
    // in the line above, msg is received from pings, and sent to pongs

    fmt.Println(<-pongs)
    // prints out the value that is received from pongs

}
