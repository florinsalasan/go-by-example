package main

import (
    "fmt"
    "os"
)

// Go offers support for string formatting in the printf tradition,
// here are some examples
type point struct {
    x, y int
}

func main() {

    p := point{1, 2}
    fmt.Printf("struct1: %v\n", p)
    // go has several printing verbs that are used to format general go
    // values, for example the above prints an instance of the point struct

    // the next verb '%+v' will include the struct's field names if applicable
    fmt.Printf("struct2: %+v\n", p)

    // the '%#v' variant prints a Go syntax representation of the value
    // ie the source code snippet that would produce the value.
    fmt.Printf("struct3: %#v\n", p)

    // To print the type of a value use %T
    fmt.Printf("type: %T\n", p)

    // formatting bools is straight forward, use %t 
    fmt.Printf("bool: %t\n", true)

    // many options for integers, %d for standard base 10 formatting
    fmt.Printf("int: %d\n", 123)

    // %b is used for binary representations of ints
    fmt.Printf("bin: %b\n", 14)

    // %c is used for character corresponding to the given int
    fmt.Printf("char: %c\n", 33)

    // %x is used for hexadecimal representations of ints
    fmt.Printf("hex: %x\n", 456)

    // also plenty of options for floats basic one is %f
    fmt.Printf("float1: %f\n", 78.9)

    // %e and %E format the float in slightly different scientific notation
    fmt.Printf("float2: %e\n", 123400000.0)
    fmt.Printf("float3: %E\n", 123400000.0)

    // basic string formatting uses %s
    fmt.Printf("str1: %s\n", "\"string\"")

    // to double quote strings use %q
    fmt.Printf("str2: %q\n", "\"string\"")

    // %x also works on strings to render it in base-16, with two output 
    // chars per byte of input
    fmt.Printf("str3: %x\n", "hex this")

    // can format a pointer with %p 
    fmt.Printf("pointer: %p\n", &p)

    // can control the width and precision of the numbers we're printing,
    // to specify the width use % followed by an int of the desired size
    // defaults to right justified and padded with spaces.
    fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

    // can specify width with floats as well, usually also restricting the 
    // decimal precision at the same time with the width.precision syntax
    fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

    // to left-justify, use the '-' flag
    fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

    // can control width in strings as well to ensure that they align in 
    // table like output, here is basic right justified width
    fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

    // can also left justify string in the same way as numbers, with '-' flag
    fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

    // only used printf and there are already so many options for verbs, those
    // are shared among multiple functions that print/format strings.
    // take for example Sprintf, which formats and returns a string rather than
    // formatting and printing to stdout like printf

    s := fmt.Sprintf("spritf: a %s", "string")
    fmt.Println(s)

    // can format and print to io.Writers other than os.Stdout when using
    // Fprintf like so:

    fmt.Fprintf(os.Stderr, "io: an %s\n", "error")


}
