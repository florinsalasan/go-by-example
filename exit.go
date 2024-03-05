package main

// use os.Exit to immediately exit with a given status

import (
    "fmt"
    "os"
)

func main() {
    // defers will not be run when using os.Exit, so this fmt.Println will 
    // never be called.
    
    defer fmt.Println("!")

    // Exit with status 3
    os.Exit(3)
}
// Unlike C Go does not use an integer return value from main to indicate 
// exit status, whenever wanting to exit with non-zero status, should
// use os.Exit. 

// Running this file the exit will be picked up by go and printed, by building
// and executing a binary can see the status in the terminal with echo $?
