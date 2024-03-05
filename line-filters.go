package main

// a line fileter is a common type of program that reads input on
// stdin, processes it, and then prints seom derived result to stdout.
// grep and sed are common line filters. Here we will make an example line
// filter that writes a capitalized version of all input text. can use this
// pattern to write other line filters after

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {

    // wrapping the unbuffered os.strdin with a buffered scanner gives us a
    // convenient scan method that advances the scanner to the next token
    // which is the next line in the default scanner.
    scanner := bufio.NewScanner(os.Stdin)

    // text returns the current token, here the next line, from the input
    for scanner.Scan() {
        ucl := strings.ToUpper(scanner.Text())

        // write out the uppercased line
        fmt.Println(ucl)

    }

    // check for errors during scan, end of file is expected and
    // not reported by scan as an error

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }

}
// to try out the filter, first make a file with a few lowercase lines
// $ echo 'hello'   > /tmp/lines
// $ echo 'filter' >> /tmp/lines

// $ cat /tmp/lines | go run line-filters.go
// HELLO
// FILTER
