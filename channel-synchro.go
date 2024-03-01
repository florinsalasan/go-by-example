package main

import (
    "fmt"
    "time"
)

// we can use channels to sync execution across goroutines, here is an example
// of using a blocking receive to wait for a goroutine  to finish.
// Once again the recommendation for syncing and waiting for async functions
// to finish WaitGroups exist and are more robust

// this is the function that will be run in a goroutine. the done channel will
// be used to notify a different goroutine that this function finished running
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    // send true to done to indicate that the function has finished running
    done <- true
}

func main() {

    done := make(chan bool, 1)
    // start a worker goroutine and passing in the channel that will notify
    // main() that worker has finished working
    go worker(done)

    // This blocks until the the worker receives a value, if removed the
    // program would exit before worker even began
    <-done

}
