package main

import "fmt"

// closing a channel indicates that no more values will be sent on it.
// this can be useful to communicate that it's completed to the 
// channel's receivers
func main() {

    // here we'll use a jobs channel to communicate work to be done from 
    // the main() goroutine to a worker goroutine, when there are no more
    // jobs for the worker we will close the jobs channel.

    jobs := make(chan int, 5)
    done := make(chan bool)

    // this will be the worker goroutine, it repeatedly receives from
    // jobs with j, more := <-jobs. in this special form of receive, the more
    // value will be false if jobs has been closed and all values in the
    // channel have already been received. This is how we will notify on
    // done when all the jobs have been worked.
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    // This sends 3 jobs to the worker over the jobs channel, then closes it
    for j := 1; j <= 3; j++ {
        // send j to jobs
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    // this is the way to await the worker using the sync approach that
    // was used in an earlier example see: channels-synchro.go
    <-done

    // reading from a closed channel succeeds immediately, returning the 
    // zero value of the underlying type. The optional second return value
    // is true if the value received was delivered by a successful send 
    // operation to the channel, or false if it was a zero value generated
    // because the channel was closed and/or empty.
    _, ok := <-jobs
    fmt.Println("received more jobs:", ok)

    // So small issue here, in the gobyexample output, the sent and received
    // print outs are alternated, when I run it locally, it sends all 3 first
    // before receiving them. Dunno if that is something that can happen or
    // if something is going wrong on my end. ran it on the go.dev playground
    // and got the same output as I do locally so I guess everything is fine.
}
