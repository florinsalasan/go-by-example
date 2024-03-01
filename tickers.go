package main

import ( 
    "fmt"
    "time"
)

// Timers are meant for doing something once in the future. tickers are for 
// doing something repeatedly at regular intervals. Here is an example 
// that will tick periodically until it is stopped.
func main() {

    // tickers use a similar mechanism to timers: a channel that is sent
    // values. Here we use select in an infinite loop to await the values
    // as they arrive every 500ms, until true is received from the channel done

    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    // just like timers, tickers can be stopped, preventing it from receiving
    // any more values on it's channel, in this case it's stopped after
    // 1600ms, meaning it should tick 3 times.
    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")

}
