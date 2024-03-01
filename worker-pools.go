package main

import (
    "fmt"
    "time"
)

// here is the worker of which we'll run multiple concurrent instances.
// the workers will receive work on the jobs channel and send the
// results to the results channel, to simulate an intensive job,
// we'll sleep 1 second per job.
func worker2(id int, jobs <-chan int, results chan<- int) {

    for j := range jobs {
        fmt.Println("worker", id, "started job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }

}

func main() {

    // to use the pool of workers we need to send them work and collect
    // results. we make jobs and results channels to do so.
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // here we start up 3 workers that are initially blocked because
    // the jobs do not exist yet.
    for w := 1; w <= 3; w++ {
        go worker2(w, jobs, results)
    }

    // here we send 5 jobs and then close the channel to indicate that
    // all the jobs have been sent.
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    // lastly we collect the results of the work ensuring that the worker
    // goroutines have finished. An alternative is once again a WaitGroup.
    for a := 1; a <= numJobs; a++ {
        <-results
    }
    // The running program should show the 5 jobs being executed by various 
    // workers and the run time should be around 2s rather than 5s had the 
    // jobs not been run concurrently.
}
