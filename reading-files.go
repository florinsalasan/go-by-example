// reading and writing files are basic tasks, here's how to accomplish
// them in go

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

// reading files requires checking most calls for errors this helper will
// streamline the process a bit
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // one of the most basic file reading tasks is reading in a file's entire
    // content into memory to be used by the program
    dat, err := os.ReadFile("/tmp/dat")
    check(err)
    fmt.Print(string(dat))

    // more often will want a bit more control over how and what parts of a 
    // file we want to read. For these more precise tasks start by opening
    // a file to get an os.File value
    f, err := os.Open("/tmp/dat")
    check(err)

    // from there we read some bytes from the beginning of the file. Allowing
    // up to 5 to be read but will check how many were actually read into mem
    b1 := make([]byte, 5)
    // seems like the file is being read into the b1 buffer
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

    _, err = f.Seek(0, 0)
    check(err)
    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: ", n2, o2)
    fmt.Printf("%v\n", string(b2[:n2]))

    // the io package provides some functions that may be helpful for file
    // reading. for example, reads like the ones above can be more robustly
    // implemented with ReadAtLeast
    _, err = f.Seek(0, 0)
    check(err)
    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    // there is no built in rewind but seek(0, 0) accomplishes the same thing
    _, err = f.Seek(0, 0)
    check(err)

    // the bufio package implements a buffered reader that may be useful both
    // for its efficiency with many small reads and because of the additional
    // reading methods it provides
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

    f.Close()
    // close the file when finished, usually this would be scheduled immediately
    // after opening the file with a defered call.

}
