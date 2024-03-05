package main

// environment variables are a universal mechanism for conveying config info
// to Unix programs, here is how to set, get, and list env variables in go

import (
    "fmt"
    "os"
    "strings"
)

func main() {

    // to set a key/value pair, use os.Setenv, to get a value, can use 
    // os.Getenv, this returns an empty string if the key is not present
    // in the environment.
    // Note these are seemingly temporary, unsure how to set them to be
    // permanent, but also don't want to right now.
    os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))

    // can use os.Environ to list all key/value pairs in the environment
    // returns a slice of strings, in the form KEY=value. Can then call
    // strings.SplitN on them to get the key and value here we'll print the keys
    fmt.Println()
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }
}

// running the program shows that we set FOO but BAR will be empty as we 
// have not given it a pair, if we go back and set a value for BAR, it
// will show up the next time it is run.
