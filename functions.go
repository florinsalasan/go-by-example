package main

import "fmt"

// here is a function that takes two ints and returns their sum
// can identify the value of each parameter individually
// requires explicit returns, won't just automatically return the 
// value of the last expression, but I don't think I've used a language
// that doesn't require explicit returns
func plus(a int, b int) int {
    return a + b
}

// When multiple consecutive parameters are of the same type, can instead 
// only declare the type on the last parameter, declaring the type for all
// preceeding parameters that were not explicitly typed
func plusPlus(a, b, c int) int {
    return a + b + c
}

// go supports multiple return values, often used in idiomatic Go to return
// result along with error values from a function. Python does similar
// but then automatically packages them together in a tuple
func vals() (int, int) {
    return 3, 7
}

// go also supports variadic functions that can be called with any number
// of trailing arguments
func sum(nums ...int) {
    // the type for nums is equal to []int, thus we can use range and len
    // to iterate over the input
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num     
    }
    fmt.Println(total)
}

// Go also supports anonymous functions which can form closures, these are
// useful for defining a function inline without having to name it.

// this function for example returns another function defined anonymously 
// inside of intSeq, and the returned function is said to close over the 
// variable i to form a closure, closures allow the function to access the 
// enclosed variable through closure copies of their values or references,
// even when the function is called outside of their scope. Won't lie, I 
// don't really understand the concept of closures, both in relation to 
// go and in general, there was a similar concept in js.
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i 
    }
}

// Recursive functions also exist in go, would have been shocked if they
// didn't more than I am that they are supported. Works just like any other
// language I've used, there is a base case and a recursive case calling
// itself until it reaches the base case. 
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n - 1)
}

// call functions just like any other language, funcName(args)
func main() {
    
    res := plus(1, 2)
    fmt.Println("1+2 = ", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 = ", res)

    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // cannot do the following and then index the return c := vals()
    // instead must use '_' to ignore values that you are not interested in
    _, c := vals()
    fmt.Println(c)

    // can call variadic functions with individual arguments
    sum(1, 2)
    sum(1, 2, 3)

    // or can pass in a slice using the func(slice...) syntax, this is 
    // used to unpack a slice and is equivalent to passing in each value
    // individually into the function call. if you have an array to pass
    // in to the variadic func, can instead do func(arr[:]...) to first
    // change the array to a slice.
    nums := []int{1, 2, 3, 4}
    sum(nums...)

    // we call intSeq assigning the result to nextInt, remember the result is
    // a function. This function will capture it's own i value, which will 
    // become updated everytime nextInt() is called.
    nextInt := intSeq()

    // call nextInt a few times to see how closure is working
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    // confirm that the state is unique to nextInt, start a new closure
    // by calling intSeq again and assigning it to a new var.
    newInts := intSeq()
    fmt.Println(newInts())

    // the recursive call 
    fmt.Println(fact(7))

    // closures can also be recursive in go, but requires the closure to be
    // declared with a typed var explicitly before defining the closure
    var fib func(n int) int

    fib = func(n int) int {
        if n < 2 {
            return n
        }
        return fib(n - 1) + fib(n - 2)
    }
    // this works only because of the declaration with a typed var 
    // before hand

    fmt.Println(fib(7))

}
