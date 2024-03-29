package main

// command line arguments are a common way to parameterize execution of programs
// for example go run hello.go uses run and hello.go arguments to go program

import (
    "fmt"
    "os"
)

func main() {

    // os.Args provides access to raw command-line arguments. Note that the
    // first value in the slice is the path to the program and os.Args[1:]
    // hold the arguments to the program
    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    // can get individual args with normal indexing
    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)

}

// to experiment with CLAs best to build a binary with go build first

