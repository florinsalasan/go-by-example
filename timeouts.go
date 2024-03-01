package main

import (
    "fmt"
    "time"
)

// timeouts are important for programs that connect to external resources or
// that otherwise need to set an upper limit on execution time. Implementing
// timeouts is done through channels and select.
func main() {

    // pretend that we're executing an external call that returns its result
    // on a channel c1 after 2s, note that the channel is buffered, so the
    // send in the goroutine is nonblocking. This is a common pattern to
    // prevent goroutine leaks in case the channel goes unread. remember
    // send in channels means a channel being sent a value to then pass along
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()

    // the select here is implementing a timeout that will proc after 1 second
    // through time.After. the select will proceed with the first receive that
    // is ready, in this first one the timeout should be the one that taken.
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }

    // since this select has a timeout that takes longer than the function
    // to run, the timeout should not occur
    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
}
