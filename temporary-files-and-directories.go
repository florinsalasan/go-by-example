package main

// throughout program execution may want to create data that isn't needed 
// after the program exits, temporary files are useful for this purpose
// since they won't pollute the file system over time

import (
    "fmt"
    "os"
    "path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // the easiest way to create a temporary file is by calling os.CreateTemp
    // it creates a file and opens it for reading and writing. We provide ""
    // as the first argument so that the function call will create the
    // file in the default location for a given OS
    f, err := os.CreateTemp("", "sample")
    check(err)

    // Display the name of the temporary file, on unix-based OSes the directory
    // should be /tmp, and the file name starts with the prefix given as the
    // second argument to os.CreateTemp and the rest is chosen automatically
    // to ensure that concurrent calls will always create diff file names

    // Interestingly when running this the temp files and directories were
    // placed in /var/folders/77/longencodednumber/createdbymeinthisfile
    fmt.Println("Temp file name:", f.Name())

    // Clean up the file after we're done. The OS is likely to clean up temp
    // files by itself after some time, but it's good practice to do
    // this explicitly.
    defer os.Remove(f.Name())

    // We can write some data to the file
    _, err = f.Write([]byte{1, 2, 3, 4})
    check(err)

    // If we intend to write many temp files, might prefer to create a temp
    // directory os.MkdirTemp's arguments are the same as CreateTemp's but
    // returns a directory name rather than an open file.
    dname, err := os.MkdirTemp("", "sampledir")
    check(err)
    fmt.Println("Temp dir name:", dname)

    defer os.RemoveAll(dname)

    // Can now synthesize temp file names by prefixing them with the temp dir
    fname := filepath.Join(dname, "file1")
    err = os.WriteFile(fname, []byte{1, 2}, 0666)
    check(err)

}
