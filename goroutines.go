package main

import (
    "fmt"
    "time"
)

// goroutines are lightweight threads of execution

// this is the function we will use to showcase the use of goroutines vs 
// calling it normally which would run it synchronously
func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {
    // this is the typical way that is running synchronously
    f("direct")

    // This is how to invoke a function in a goroutine. This new goroutine
    // will execute concurrently with the calling one above
    go f("goroutine")

    // can also start goroutines on anonymous function calls
    go func(msg string) {
        fmt.Println(msg)
    }("going")

    // the function calls are in two separate goroutines 
    // running asynchronously. We can use a naive approach of waiting them 
    // to finish by using time.Sleep, a more robust approach uses 
    // a WaitGroup
    time.Sleep(time.Second)
    fmt.Println("done")
}

// when we run this file, we see the output of the blocking call first,
// then the output of the two goroutines. Depending on the time taken for
// parts of the goroutines to run, the outputs may be woven together rather
// than showing up separately because they are being run concurrently
