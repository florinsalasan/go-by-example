package main

import (
    "fmt"
    "maps"
)

func main() {
    
    // syntax for making an empty map is: "make(map[key-type]value-type)"
    m := make(map[string]int)

    // Can then set key-value pairs by using "name-of-map[key] = value" syntax
    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)
    // printing a map will show all key-value pairs

    // can access a value using name-of-map[key]
    v1 := m["k1"]
    fmt.Println("v1:", v1)

    // if a key does not exist, the zero value for the value type will be 
    // returned instead
    v3 := m["k3"]
    fmt.Println("v3:", v3)

    // maps have a built in len that returns the number of key-value pairs
    fmt.Println("len:", len(m))

    // the delete builtin removes key-value pairs when called on a map
    // done with the following syntax: "delete(map-name, key-name)"
    delete(m, "k2")
    fmt.Println("map:", m)

    // clear builtin will remove all key-value pairs from the map
    clear(m)
    fmt.Println("map:", m)

    // can have an optional second return value to indicate whether a key is 
    // present in a map, useful since go doesn't prevent you from grabbing a
    // value from a key that doesn't exist. If you do not need the actual 
    // value from a check to see if the key exists, can use '_' as the blank
    // identifier. In most cases I'd imagine you'd want to use both, and then
    // before using the value, checking the present variable to ensure that
    // the pairing exists.
    _, present := m["k2"]
    fmt.Println("prs:", present)

    // can also declare and initialize a map in one line similar to slices
    // and arrays
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

    // just like the Slices package, there are plenty of utility functions
    // included in the Maps package
    n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }
}
