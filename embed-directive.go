package main

// go:embed is a compiler directive that allows programs to include 
// arbitrary files and folders in the go binary at build time.
// more info at https://pkg.go.dev/embed

import (
    "embed"
)

// import the embed package, if you don't use any exported identifiers from
// the package, can do blank import with _ "embed"
//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte
// embed directives accept paths relative to the directory containing the go
// source file, this directive embeds the contents of the file into the 
// string variable immediately following it, can also use []byte

// can embed multiple files or even folders with wildcards, this uses a
// variable of the embed.FS type, which implements a simple virtual file system

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

    // print the contents of single_file.txt
    print(fileString)
    print(string(fileByte))

    // retrieve some files from the embedded folder
    content1, _ := folder.ReadFile("folder/file1.hash")
    print(string(content1))

    content2, _ := folder.ReadFile("folder/file2.hash")
    print(string(content2))
}

// Use the following to test out this example
// mkdir -p folder
// $ echo "hello go" > folder/single_file.txt
// $ echo "123" > folder/file1.hash
// $ echo "456" > folder/file2.hash

// $ go run embed-directive.go
// hello go
// hello go
// 123
// 456
