package main

import "fmt"

// in go arrays are a numbered sequence of elements of a specific length
func main() {

    var a [5]int
    // all 5 ints are zero valued
    fmt.Println("emp:", a)

    a[4] = 100
    // last int is set to 100, can use Array[index] to get the value at
    // position index, just like many other languages
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))
    // len(Array) should be self explanatory

    b := [5]int{1, 2, 3, 4, 5}
    // use the above to declare and initialize the array with specific values
    // unlike python, values are initialized within {} 
    fmt.Println("dcl:", b)

    var twoD [2][3]int
    // by default arrays are 1D, can compose them to make multi dimensional 
    // arrays just like many other languages
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
