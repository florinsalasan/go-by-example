package main

import "fmt"

func main() {
    
    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)
    
    var d = true // go can infer the type of the variable
    fmt.Println(d)

    var e int // variables declared without initializing are zero values
    fmt.Println(e) // prints 0

    f := "apple" // short form for declaring and initializing a var
    fmt.Println(f)
}
