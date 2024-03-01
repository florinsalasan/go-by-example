package main

import (
    "fmt"
    "time"
)

// go's select allows waiting on multiple channel operations.
// combining goroutines and channels with select is a powerful 
// feature of the language
func main() {

    // for our example, will select across two channels
    c1 := make(chan string)
    c2 := make(chan string)

    // each channel will receive a value after some time, to simulate remote
    // procedure call operations, that are blocking by default, that are
    // executing in concurrent goroutines.
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    // We'll use select to await both of the values simultaneously, printing
    // each one as it arrives. This works because we're looping over the select
    // statement and since each channel only sends/receives once. If they
    // were sent to multiple times, it might not get both cases.
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
    // The program prints out, one, two in order and the program runs in
    // about ~2 seconds since the sleep calls occur concurrently rather than
    // one after the other.
}
