package main

import (
    "errors"
    "fmt"
)

// in go it is idiomatic to communicate errors with an explicit, separate 
// return value. This is unlike the exceptions used in many other languages
// like Java and unlike the overloaded single return that is typical of c

// This approach makes it easier to see which function returns errors and
// to handle them in ways similar to other non-error tasks.
func f1(arg int) (int, error) {
    // by convention the error is the last return value and have type error
    if arg == 42 {
        return -1, errors.New("Can't work with 42")
        // errors.New contructs a basic error value with the passed in
        // error message.
    }
    return arg + 3, nil
    // returning nil in the error position indicates that there was no error.
}

type argError struct {
    arg int
    prob string
}
// It's possible to use custom types as errors by implementing the Error method
// on them. Here is an example that is customized to represent an argument error

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {
        // here we use &argError syntax to build a new struct, passing in
        // values for the arg and the problem
        return -1, &argError{arg, "can't work with 42!!!"}
    }
    return arg + 3, nil
}

func main() {
    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked", r)
        }
    }
    _, e := f2(42)
    // if wanting to programmatically use the data in a custom error,
    // need to get the error as an instance of the custom error type via
    // type assertion.
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}

