package main

import "fmt"

func main() {

    i := 1
    for i <= 3 { // all loop types are for loops in go, this one for example
        // is a while loop 
        fmt.Println(i)
        i = i + 1
    }

    // This next one is more of a classic style for loop, with initial,
    // condition, and after paramters for the loop
    for j := 0; j < 3; j++ {
        fmt.Println(j)
    }

    // can also use range like python for loops in the following way:
    for i := range 3 {
        fmt.Println("range", i)
    }

    // can also set up an infinite loop if you want. only way it will break
    // is when you return from the loop or break out of it
    for {
        fmt.Println("loop")
        break // this should only print 'loop' once but could be dangerous
    }

    for n := range 6 {
        if n % 2 == 0 {
            continue // can choose to skip to the next value under certain
            // conditions with continues just like other language
        }
        fmt.Println(n)
    }

}
