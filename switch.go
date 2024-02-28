package main

import (
    "fmt"
    "time"
)

// switch statements express conditionals across many branches just
// like in other languages that have switch statements
func main() {
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
        case 1:
            fmt.Println("one")
        case 2:
            fmt.Println("two")
        case 3:
            fmt.Println("three")
    }

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday: // can use commas to separate multiple
        // expressions in the same case statement.
        fmt.Println("It's the weekend")
    default:
        // default is what is returned when none of the other cases are true
        fmt.Println("It's a weekday")
    }

    t := time.Now()
    // can use switch statements without an expression as an alternate if/else
    // block of statements. the case expressions can also be non-constants
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        // a type switch compares types rather than values, can do this to 
        // discover the type of an interface value, variable t will have
        // the type corresponding to the type of the value passed into
        // whatAmI
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }    
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}
