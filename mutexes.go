package main

import (
    "fmt"
    "sync"
)

// just finished handling simple state with atomic operations, for more
// complex state mutex can be used to safely access data across
// multiple goroutines

// The container willl hold a map of counters, since we want to update it
// concurrently from multiple goroutines, we add a mutex to sync access. 
// Note that mutexes should not be copied, so if the struct is passed around
// it should be done so using pointers.
type Container struct {
    mu sync.Mutex
    counters map[string]int
}

func (c *Container) inc(name string) {

    // we lock the mutex before accessing counters, and then the container
    // is unlocked at the end of the function with the defer statement

    // it locks the mutex, to ensure that other goroutines are not accessing
    // the struct that is being modified, does whatever changes needed from
    // the function, and then unlocks it to once again be modifiable by other
    // goroutines. This ensures safety when accessing data across many goroutines
    c.mu.Lock()
    defer c.mu.Unlock()
    c.counters[name]++

    // also note that mutexes aren't always the answer to syncing up, 
    // if used improperly could easily be a bottleneck in the program

}

func main(){

    c := Container {
        // note that a mutex isn't explicitly defined as the zero value
        // is usable as is, do not need to initialize it
        
        counters: map[string]int{"a": 0, "b":0},

    }

    var wg sync.WaitGroup

    // the function increments a given counter in a loop
    doIncrement := func(name string, n int) {
        for i := 0; i < n; i++ {
            // function is defined to only work on the existing
            // container, rather than taking a container in
            // as a pointer 
            c.inc(name)
        }
        wg.Done()
    }

    // run 3 goroutines concurrently which all access the same container, 
    // and two of which access the same counter.
    wg.Add(3)
    go doIncrement("a", 10000)
    go doIncrement("a", 10000)
    go doIncrement("b", 10000)

    // wait on the goroutines to finish
    wg.Wait()
    fmt.Println(c.counters)

}
