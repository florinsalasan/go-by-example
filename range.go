package main

import "fmt"

func main() {

    nums := []int{2, 3, 4}
    sum := 0
    // Can use range in a similar way to python range, to iterate over elements
    // in a lot of data structures, it has two values for arrays and slices,
    // returning the index and the value in that order, if you don't need the 
    // info about one of them can use '_' again to ignore it
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // range can be used to iterate over key-value pairs
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    // can also iterate just over the keys of a map
    for k := range kvs {
        fmt.Println("key:", k)
    }

    // can also use range on strings, where the first value is the starting
    // byte index of the 'rune', think char in other languages, and then the 
    // second value is the rune itself, represented as the ascii value of the 
    // character. so 'g' would be 103
    for i, c := range "go" {
        fmt.Println(i, c)
    }

}
