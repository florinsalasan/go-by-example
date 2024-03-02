package main

import (
    "fmt"
    "time"
)

// rate limiting is important for resource utilization and maintaining
// quality of service. Go supports this with goroutines, channels, and
// tickers, all of which have been covered.

func main() {

    // firstly some basic rate limiting. Suppose we want to limit 
    // handling of some incoming requests, we can serve all of the requests
    // off of one channel.
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    // reminder that time.Tick is a channel as the Ticker struct
    // contains one attribute 'C' that is the channel of type time
    // also using time.Tick isn't exactly safe and cannot be garbage 
    // collected if there is no way to shut down the Ticker, better to 
    // use time.NewTicker which is what time.Tick is a convenience wrapper
    // for anyways.
    limiter := time.Tick(200 * time.Millisecond)
    // the limiter channel will receive a value every 200ms, and will 
    // act as the regulator in the rate limiting setup.

    for req := range requests {
        // wait on limiter to receive a value, which we set to happen every
        // 200ms, once received it will stop blocking
        <-limiter
        // by blocking on a receive from the limiter, we rate limit these
        // 'requests' to one every 200ms
        fmt.Println("regular request", req, time.Now())
    }

    // We may also want to allow a certain number of requests to be able to 
    // come in at once in a 'burst' while also preserving the overall rate 
    // limit, we can accomplish this by buffering the channel we use as the 
    // limiter, this one allows bursts of up to 3 events
    burstyLimiter := make(chan time.Time, 3)

    // we fill up the channel to represent the allowed bursts of 'requests'
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    // we start a goroutine that attempts to add a new value to the burstylimiter
    // every 200ms until it reaches it's limit of 3
    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()

    // We then simulate another 5 requests, this time on the burstyLimiter
    // Based on the setup the first 3 requests should be able to come in
    // without being limited, and then the remaining two should be held
    // to the same 200ms rate limit as before.
    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)

    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("bursty request", req, time.Now())
    }
    // After running we see exactly what we expected, the first limiter spaces
    // the requests evenly every 200ms, the burst allows for the first 3 
    // requests to come in almost simultaneously and then every 200ms for the 
    // other 2. The main question I have is how is the overall rate limit
    // set for the burst? and how would I be able to allow for a second 
    // burst if it was within the rate limit set for the program overall?
}
