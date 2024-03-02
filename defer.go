package main

import (
    "fmt"
    "os"
)
// defer is used to ensure that a function call is performed later in
// a program's execution, usually for purposes of cleanup, defer is often
// used where things like 'ensure' and 'finally' are used in other languages.

func main() {
    // suppose we want to create a file, write to it, then close it when
    // we're finished, here is how defer can be used.

    f := createFile("/tmp/defer.txt")
    // right after we create the file we defer closing it which means that
    // closeFile will only be executed at the end of the enclosing function
    // in this case main(), after writeFile finishes

    // side note, would write file not be blocking here? could we not just
    // call writeFile before closeFile and everything would be fine?
    // maybe I'm too used to python file usage
    defer closeFile(f)
    writeFile(f)

}

func createFile(p string) *os.File {

    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f

}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}

// important to check for errors when closing a file even in deferred funcs.
// I'd imagine to ensure that the file isn't corrupted? not explained on site.
func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()

    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}
