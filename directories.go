package main

// go has several functions for working with directories
import (
    "fmt"
    "io/fs"
    "os"
    "path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // create a new sub-directory in the current working directory
    err := os.Mkdir("subdir", 0755) // make a new directory with 0755 perms
    check(err)

    // good practice to defer the removal of temporary sub-directories after
    // creating them. os.RemoveAll will delete a whole directory tree, kinda
    // like how rm -rf works
    defer os.RemoveAll("subdir")

    // Helper function to create a new empty file:
    createEmptyFile := func(name string) {
        d := []byte("")
        check(os.WriteFile(name, d, 0644))
    }

    createEmptyFile("subdir/file1")

    // can create a hierarchy of directories, including parents with 
    // MkdirAll, similar to how mkdir -p works (mkdir -p grandparent/parent/file
    // creates the parent folder if it does not exist yet)
    err = os.MkdirAll("subdir/parent/child", 0755)
    check(err)

    createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

    // ReadDir lists directory contents, returning a slice of os.DirEntry objs
    c, err := os.ReadDir("subdir/parent")
    check(err)

    fmt.Println("Listing subdir/parent")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    // Chdir lets us change the current working directory, similarly to cd
    err = os.Chdir("subdir/parent/child")
    check(err)

    // Now will see the contents of 'subdir/parent/child' when listing the 
    // current directory
    c, err = os.ReadDir(".")
    check(err)

    fmt.Println("Listing subdir/parent/child")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    // cd back to where we started:
    err = os.Chdir("../../..")
    check(err)

    // can also visit a directory recursively, including all its sub-directories
    // by using WalkDir, which accepts a callback function to handle each
    // file or directory visited.
    fmt.Println("visiting sub-directory")
    err = filepath.WalkDir("subdir", visit)

}

func visit(path string, d fs.DirEntry, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", path, d.IsDir())
    return nil
}
