package main

import "fmt"

// we know range and for can iterate over various data structures, we can
// also iterate over values received from a channel.

func main() {

    // we'll iterate over 2 values in the queue channel
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    // this range iterates over each element as it's received from queue
    // because we closed the channel before iterating, the iteration will
    // terminate after receiving the 2 elements.
    for elem := range queue {
        fmt.Println(elem)
    }
    // This example shows that it's possible to close a channel before
    // receiving all of the values and emptying out the channel.
}
