package main

import (
    "fmt"
    "slices"
)

// slices are similar to arrays but more flexible and just more powerful
// think of the yearbook meme, daniel and cooler daniel

func main() {

    var s []string
    // note how you don't need to pass in a length when declaring a slice,
    // unlike arrays, this might actually be the one difference in declaring
    // a slice vs an array
    fmt.Println("uninit:", s, s == nil, len(s) == 0)
    // slices are typed by the elements that they contain thus an uninitialized
    // slice will be equal to nil, and have a len of 0.

    s = make([]string, 3)
    // to make an empty slice use make, by default capacity should equal len,
    // if you plan on increasing the size later on can instead define the
    // capacity explicitly as the last parameter to make
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
    // slices have the values, len, and capacity.
    
    // can set and get just like arrays
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    // appending is a bit strange compared to append in python or push in js
    // returns a slice and takes in the slice appending to, along with the 
    // values being appended, append does not work on arrays, I think
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // slices can also be copied where we can create an empty slice of the 
    // len of s, then copy values from s into the new slice
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // slices can then be sliced with indices just like python lists can be
    // the first value is included, and the second value is not included in
    // the new slice, so getting items 3-5 from an array would be a[3:6]
    // just like python can leave one of the parameters empty to get from start
    // second index - 1, or from index to end as a[:4] or a[2:] can even have
    // a[:] if wanted, but don't think that is useful tbh.
    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    // can declare and initialize all in one line just like arrays.
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)
    fmt.Printf("t is of type %T\n", t)

    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d:", twoD)
    // While slices are not the same as arrays, when printed out they are 
    // rendered in a similar fashion.
}

