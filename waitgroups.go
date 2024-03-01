package main

import (

    "fmt"
    "sync"
    "time"

)

// to wait for multiple goroutines to finish, waitgroups are the recommended 
// way of doing this.

// this is the worker that we will use in the example, once again
// sleeping for a second to simulate a compute heavy task
func worker3(id int) {
    fmt.Printf("Worker %d starting\n", id)
    
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {

    // this waitgroup is used to wait for all the goroutines launched here to 
    // finish. if a waitgroup is explicitly passed into functions, it should
    // be done by a pointer.
    var wg sync.WaitGroup

    // this launches 5 workers and increments the waitgroup counter for each
    for i := 1; i <= 5; i++ {

        wg.Add(1)

        // wrap the worker call in a closure that will make sure to tell the 
        // waitgroup that the specific worker is done. This way the worker 
        // does not need to be aware of the concurrency primitives involved
        // in its execution
        go func() {
            defer wg.Done()
            worker3(i)
        }()

    }

    // Blocks until the waitgroup counter goes back to 0; and all the workers
    // have notified that they are done running
    wg.Wait()

    // the one downside is that this approach has no straightforward way to 
    // propagate errors from the workers. For more advanced use cases
    // there is an errgroup package that can help

    // the order workers start and finish can be different for each time 
    // the program is run.

}
