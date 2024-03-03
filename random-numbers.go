package main

import (
    "fmt"
    "math/rand/v2"
)

// math/rand/v2 package provides pseudorandom number generation

func main() {

    // for example rand.IntN returns a random int n, 0 <= n < 100
    fmt.Print(rand.IntN(100), ",")
    fmt.Print(rand.IntN(100))
    fmt.Println()

    // rand.float64 returns a float64 f, 0.0 <= f < 1.0
    fmt.Println(rand.Float64())

    // This can be used to generate rand floates in other ranges, ie 
    // between 5.0 <= f < 10.0
    fmt.Print((rand.Float64()*5) + 5, ",")
    fmt.Print((rand.Float64()*5) + 5)
    fmt.Println()

    // can also generate a seed for a more consistent value
    // first create a new rand.Source and pass it into the New constructor
    // NewPCG creates a new PCG source that reuires a seed of two uint64 nums
    s2 := rand.NewPCG(42, 1024)
    r2 := rand.New(s2)
    fmt.Print(r2.IntN(100), ",")
    fmt.Print(r2.IntN(100))
    fmt.Println()

    s3 := rand.NewPCG(42, 1024)
    r3 := rand.New(s3)
    fmt.Print(r3.IntN(100), ",")
    fmt.Print(r3.IntN(100))
    fmt.Println()
}
