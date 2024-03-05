package main

// Sometimes want go to intelligently handle Unix signals, ie might want a 
// server to gracefully shutdown when it receives a SIGTERM, or a CLI tool
// to stop processing input if it receives a SIGINT. Here is how to handle
// signals in Go with channels

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {

    // Go signal notification works by sending os.Signal values on a channel
    // will create a channel to receive these notifications note that the 
    // channel should be buffered.
    sigs := make(chan os.Signal, 1)

    // signal.Notify registers the given channel to receive notifications of 
    // the specified signals.
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    // Could receive from sigs here in the main function, but to see how
    // this could also be done in a separate goroutine, to demonstrate a more
    // realistic scenario of graceful shutdown
    done := make(chan bool, 1)

    // the following goroutine will execute a blocking receive for signals
    // when it gets one it'll print it out and notify the program that it 
    // can finish
    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    fmt.Println("awaiting signal")
    // The program waits here until it gets the expected signal from the 
    // goroutine sending a value on done, and then exit
    <-done
    fmt.Println("exiting")

}

// when running the program it will block waiting for a signal, by hitting
// ctrl+c can send a SIGINT signal causing the program to interrupt and 
// then exit
