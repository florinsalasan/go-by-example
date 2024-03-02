package main

import (
    "fmt"
    "slices"
)

// Go's slices package implements sorting for builtins and user-defined 
// types, here are some examples of sorting builtins.

func main() {

    // sorting functions are generic and work for any ordered builtin
    // for a list of ordered types in go see: https://pkg.go.dev/cmp#Ordered
    strs := []string{"c", "a", "b"}
    slices.Sort(strs)
    fmt.Println("strings:", strs)

    ints := []int{7, 2, 4}
    slices.Sort(ints)
    fmt.Println("Ints: ", ints)

    // can also use IsSorted method to check if a slice is already sorted
    s := slices.IsSorted(ints)
    fmt.Println("Sorted: ", s)
}
