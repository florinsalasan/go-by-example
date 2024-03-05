package main

// the filepath package provides functions to parse and construct file paths 
// in a way that is portable between operating systems, important since
// Linux and Windows use opposite '/' and '\' for directory separators

import (
    "fmt"
    "path/filepath"
    "strings"
)

func main() {

    // join should be used to construct paths in a portable way. takes any
    // number of arguments and constructs a hierarchical path from them
    p := filepath.Join("dir1", "dir2", "filename")
    fmt.Println("p:", p)

    // Should always use join instead of concatenating with the slashes 
    // manually. In addition to making it portable, join also normalizes paths
    // by removing superfluous separators and directory changes
    fmt.Println(filepath.Join("dir1//", "filename"))
    fmt.Println(filepath.Join("dir1/../dir1", "filename"))

    // Dir and Base can be used to split a path to the directory and the file
    // alternatively Split will return both in the same call.
    fmt.Println("Dir(p):", filepath.Dir(p))
    fmt.Println("Base(p):", filepath.Base(p))

    // can check if a path is absolute
    fmt.Println(filepath.IsAbs("dir/file"))
    fmt.Println(filepath.IsAbs("/dir/file"))

    filename := "config.json"

    // some file names have extensions following a dot. The extension can be
    // split from the filename with Ext
    ext := filepath.Ext(filename)
    fmt.Println(ext)

    // to find the file's name with the extension removed use 
    // strings.TrimSuffix
    fmt.Println(strings.TrimSuffix(filename, ext))

    // Rel finds a relative path between a base and a target, it returns an
    // error if the target cannot be made relative to base
    rel, err := filepath.Rel("a/b", "a/b/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)

    rel, err = filepath.Rel("a/b", "a/c/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
}

