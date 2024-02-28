package main

import "fmt"

func main() {

    if 7%2 == 0 { // do not need to have the condition within parentheses
    // do need brackets for the resulting statement flow
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    if 8%4 == 0 { // do not need a matching else statement, can just have 
        // a conditional that tweaks something and then the program
        // continues after regardless or can use this for early returns
        fmt.Println("8 is divisible by 4")
    }

    if 8%2 == 0 || 7%2 == 0 { // or is || like c, and is &&
        fmt.Println("7 or 8 or both are even")
    }

    // can precede conditionals with a statement, the variables declared in
    // this statement are available in the if-else branches
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

    // Note go does not have a ternary operator, which frankly I prefer
    // because keeping if-else is intuitive rather than ternaries which
    // may not behave the same across languages that support it.
}
