package main

import (
    "bufio"
    "fmt"
    "os"
)

// writing files in go use a similar workflow to reading them

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // to begin we'll dump a string (or just bytes) into a file
    d1 := []byte("hello\ngo\n")
    err := os.WriteFile("/tmp/dat1", d1, 0644) // 0644 are the file permissions
    check(err)

    // for some more granular writes, open a file for writing.
    f, err := os.Create("/tmp/dat2")
    check(err)
    // idiomatic to defer a close immediately after opening a file and 
    // ensuring there wasn't an error
    defer f.Close()

    // can write byte slices as expected
    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    // a writestring method is also available
    n3, err := f.WriteString("writes\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n3)

    // issue a Sync to flush writes to stable storage.
    f.Sync()

    // bufio provides buffered writers in addition to the buffered readers
    // we've used before. 
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n4)

    // use Flush to ensure all buffered operations have been applied
    // to the underlying writer
    w.Flush()

}


