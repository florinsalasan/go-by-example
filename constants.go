package main

import (
    "fmt"
    "math"
)

// const supports character, string, boolean, and numeric values

const s string = "constant" // can use a const any time a var can be used

func main() {
    fmt.Println(s)

    const n = 500000000

    const d = 3e20 / n // const expressions perform math with arbitrary precision
    fmt.Println(d)

    fmt.Println(int64(d)) // const d would have no type until given one such 
    // as by explicit conversion

    fmt.Println(math.Sin(n)) // can also be given a type by using it in a 
    // context that requires one like variable assignment or function call,
    // math.Sin for example requires a float64
}
