package main

import "fmt"

// go allows for programs to recover from a panic, by using the builtin
// recover function. A recover can stop a panic from aborting the 
// program and let it continue with execution instead. 

// This would be useful in a server that doesn't want to crash if a client
// connection exhibits a critical error, instead it would probably be far
// better for that connection to be closed instead, and the server continues
// working for the other clients. This is waht go's net/http does by default
// for HTTP servers.
func mayPanic() {
    // make a function that just panics.
    panic("a new problem")
}

func main() {

    defer func() {
        // recover must be called within a defered function, when the enclosing
        // function panics the defer will activate and the recover call 
        // within will catch the panic
        if r := recover(); r != nil {
            fmt.Println("Recovered. Error:\n", r)
            // the return value of recover is the error that was raised
            // when the program panicked.
        }
    }()

    mayPanic()

    // this print will never run because the program panics and then resumes
    // within the defered function that contains the recover
    fmt.Println("after mayPanic()")

}
