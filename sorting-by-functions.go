package main

import (
    "fmt"
    "cmp"
    "slices"
)
// Sometimes want to sort collection by a custom parameter. For example
// sorting strings by len rather than alphabetically, here are some examples
// of custom sorts and how to implement them

func main() {
    fruits := []string{"peach", "banana", "kiwi"}

    // comparison function is implemented for string lengths, cmp.Compare
    // makes this quite trivial.
    lenCmp := func(a, b string) int {
        return cmp.Compare(len(a), len(b))
    }

    // then call slices.SortFunc with the slice we are sorting and the custom
    // comparison function to sort the fruits by the name length
    slices.SortFunc(fruits, lenCmp)
    fmt.Println(fruits)

    // same technique can be used to sort a slice of values that are not
    // built in types
    type Person struct {
        name string
        age int
    }

    people := []Person{
            Person{name: "Jax", age: 37},
            Person{name: "TJ", age: 25},
            Person{name: "Alex", age: 72},
    }

    // sort people by age, implemented with cmp.Compare again
    slices.SortFunc(people,
            func(a, b Person) int {
                return cmp.Compare(a.age, b.age)
        })
    // if Person structs or whatever custom type you are sorting is too large
    // and performance isn't good enough, can instead sort on a slice of
    // pointers to the large struct instead, adjusting the sort function
    // as needed.
    fmt.Println(people)
}

