package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

// in the mutexes example we used explicit locking with mutexes to sync
// access to shared state across goroutines, another option is to use
// the built-in sync features of goroutines and channels to achieve the 
// same result. This channel-based approach aligns with Go's ideas of sharing
// memory by comminucating and having each piece of data owned by exactly 
// 1 goroutine


// In this example the state will be owned by one goroutine, this will 
// guarantee that the data is never corrupted with concurrent access. in 
// order to read or write to the state, other goroutines will send messages
// to the owning goroutine and receive corresponding replies. The readOp and 
// writeOp structs will encapsulate the requests and a way for the state
// goroutine to respond.
type readOp struct {

    key int
    resp chan int

}

type writeOp struct {

    key int
    val int
    resp chan bool

}

func main() {

    // these will count the number of operations we perform throughout
    // the program
    var readOps uint64
    var writeOps uint64

    // these are the channels that will be used by the other goroutines
    // to issue read and write requests
    reads := make(chan readOp)
    writes := make(chan writeOp)


    // this is the goroutine that will handle the state
    // The state is still a map just like in the mutex example, but it is now
    // private to this goroutine. The routine will repeatedly select on the 
    // reads and writes channels responding to requests as they come in. A 
    // response is then executed by first performing whatever operation was 
    // requested and then sending a value on the response channel resp to 
    // indicate success and the value requested in a read operation.
    go func() {

        var state = make(map[int]int)
        for {

            select {
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }

        }

    }()

    // this starts 100 goroutines that issue reads to the state routine via
    // the reads channel. Each read requires constructing a readOp and then
    // sending it over the reads channel, finally receiving the result over
    // the resp channel.
    for r := 0; r < 100; r++ {

        go func () {
            for {
                read := readOp{
                    key: rand.Intn(5),
                    resp: make(chan int)}
                reads <- read
                <-read.resp
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // this starts 10 goroutines responsible for writing to the state
    for w := 0; w < 10; w++ {

        go func () {
            for {
                write := writeOp{
                    key: rand.Intn(5),
                    val: rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write 
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second)

    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writereadOps:", writeOpsFinal)

    // this method is a bit more involved than the mutex based approach, 
    // however it can be useful when managing multiple mutexes would be 
    // error-prone or when multiple channels are involved. 

    // Use whichever one makes more sense for the situation that is trying to 
    // be solved.
}
