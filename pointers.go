package main

import "fmt"

// go supports pointers allowing us to pass references to values and records
// within the program

// will the diff between using pointers and values with two different functions
// zeroval and zeroptr. zeroval gets an int passed in as a value, zeroval
// will get a copy of ival that is distinct from the one in the calling func
func zeroval(ival int) {
    ival = 0
}

// zeroptr instead takes in a *int parameter, or an int pointer. the *iptr in
// the function will dereference the pointer from its memory address to the 
// current value at that address. Assigning a value to a dereferenced ptr
// changes the value at the referenced address.
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {

    i := 1
    fmt.Println("initial: ", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    // &i syntax passes in the memory address of i, aka a pointer to i
    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    // can also print out the pointer
    fmt.Println("pointer:", &i)

    // Note how zeroval doesn't change the i in main, but zeroptr will
    // because it has access to the memory address where i is stored
    // and thus can change it, unlike zeroval which only gets a copy of i,
    // that is distinct from the one we want to change.
}
