package main

import (
    "fmt"
    "time"
)

// sometimes we want to execute go code at a specific point in the future
// or in a given interval. Go has a built in timer and ticker feature that
// helps us accomplish that. First looking at timers.
func main() {

    timer1 := time.NewTimer(2 * time.Second)
    // timers represent a single event in the future. the time is given the 
    // amount of time you want it to wait for, it provides a channel that
    // will be notified at that time. This example timer is set to 
    // wait 2 seconds.

    <-timer1.C
    fmt.Println("Timer 1 fired")
    // This line blocks on the timer's channel C until it sends a value
    // indicating that the timer fired. Channel C is the channel contained in
    // Timer structs.

    // if we just wanted to wait, could have used time.Sleep. However a timer
    // is useful because the timer can be cancelled before it fires. Here 
    // is an example of that.
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }

    // gives timer2 enough time to fire, if it was going to, to show that 
    // it was stopped instead.
    time.Sleep(2 * time.Second)

    // the first timer should fire 2s after running the program, but the
    // second should never fire as it should have been stopped

}
