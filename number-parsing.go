package main

// parsing numbers from strings is basic yet common, so this is how to 
// do it in go

import (
    "fmt"
    "strconv"
)

// the built in package strconv provides the number parsing
func main() {

    // with parse float the second parameter tells how many 
    // bits of precision we want
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    // parse int takes in a string, the 0 tells it to infer the base
    // from the string, and the 64 tells it to fit the result within 64 bits
    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

    // parse int can recognize hexadecimal numbers
    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

    // parseUint exists
    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

    // Atoi is a convenience function for basic base-10 int parsing
    k, _ := strconv.Atoi("135")
    fmt.Println(k)

    // parse functions return errors on bad inputs
    _, e := strconv.Atoi("wat")
    fmt.Println(e)

}
