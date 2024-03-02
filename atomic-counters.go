package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

// So the main mechanism for managing state in Go is communcation over channels
// we saw this when working with worker pools, there are a few other methods
// to manage state with. Here we will use sync/atomic package for atomic 
// counters accessible by multiple goroutines

func main() {

    // We use an atomic integer type to represent our (always-positive) counter
    var ops atomic.Uint64
    //var ops uint64

    // A waitgroup will help us wait for all goroutines to finish their work
    var wg sync.WaitGroup

    // Next we'll start 50 goroutines that will each increment the atomic
    // counter 1000 times
    for i := 0; i < 50; i++ {

        wg.Add(1)

        go func() {
            for c := 0; c < 1000; c++ {
                // to increment the counter use the built in Add()
                //ops++
                ops.Add(1)
                // by lowering the times incremented to 10, can easily
                // see the interweaving increments from different goroutines
                // fmt.Println(i)
            }
            wg.Done()
        }()
    }

    // we wait for all of the go routines to end
    wg.Wait()

    // no goroutines are writing to ops anymore, but using Load() it is safe to
    // atomically read a value even while other goroutines are (atomically)
    // updating it.
    // fmt.Println("ops:", ops)

    fmt.Println("ops:", ops.Load())
    // expect to get a counter of exactly 50k, had we used non atomic integer
    // and incremented with ops++ there would be no guarantee of getting 50k,
    // it would likely not be consistent across runs, the goroutines would 
    // interfere with one another, and we could even run into data race
    // failures when running with the -race flag
}
