package main

import "os"
// panics typically mean something went unexpectedly wrong. mostly used to
// fail fast on errors that shouldn't occur during normal operation or 
// that we failed to prepare to handle gracefully.

func main() {

    // going forward panics will be used to check for unexpected errors, 
    // this is the only program designed to panic.
    panic("a problem")

    // a common use of panic is to abort if a function returns an error 
    // value that we don't know how/don't want to handle. Here is an example
    // of panicking if we get an error when trying to create a new file,
    // although my ide is saying this is unreachable code.

    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }

    // when running this, the program will panic as expected and print 
    // an error message and goroutine traces, then exit with non-zero status

    // This means that the only way for the second panic to ever go off in 
    // this file requires removing the first panic.

    // In go it is idiomatic to use error-indicating return values wherever
    // possible in place of exceptions for handling errors

}
