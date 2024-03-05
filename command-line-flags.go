package main

// command line flags are a common way to specify options for command line
// programs. ie in wc -l, the -l is a command line flag

import (
    "flag"
    "fmt"
)

// go has a flag package supporting basic command line flag parsing. will use
// this package to implement our example program

func main() {

    // Basic flag declarations are available for string, integer, and bool
    // options, here we declare a string flag word, with a default value foo
    // and a short description. This flag.String function returns a string ptr
    // not a value, and we'll use this pointer further on
    wordPtr := flag.String("word", "foo", "a string")

    // declare numb and fork flags, using similar approach to the str flag
    numbPtr := flag.Int("numb", 42, "an int")
    forkPtr := flag.Bool("fork", false, "a bool")

    // also possible to declare an option that uses an existing var declared
    // somewhere else in the program. note that this requires passing in a 
    // pointer to the flag declaration function
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    // Once the flags are all declared can call flag.Parse() to execute
    // the command-line parsing
    flag.Parse()

    // Here we'll print out the parsed options and any remaining positional
    // arguments in the tail:. Note that we need to dereference the pointers
    // with '*' to get the actual option values instead of the
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *forkPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}

// To experiment with the command-line flags, it's best to compile and then
// run the binary after

// First try to give it values for all flags
// Then note that if flags are omitted they will take their default values
// Trailing positional arguments can be provided after any flags, if a flag
//    is provided after positional arguments start being provided, the flag
//    will be treated as a positional argument and won't be parsed.
// Use -h or --help flags to get the automatically generated help text
//    for the command-line program
// if a flag is provided that wasn't specified in the program, the program
//    print an error message and show the help text too.
