package main

// go std lib comes with support for http clients and servers in the net/http
// package. In this example will use it to issue simple HTTP requests.

import (
    "bufio"
    "fmt"
    "net/http"
)

func main() {

    // Issue an HTTP GET request to a server, http.Get is a convenient shortcut
    // around creating an http.Client object and then calling it's get method
    // it uses the http.DefaultClient object which has resonable default settings
    resp, err := http.Get("https://gobyexample.com")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    // print the HTTP response status
    fmt.Println("Response status:", resp.Status)

    // Print the first 5 lines of the response body
    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
// Run the client.
