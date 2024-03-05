package main

// Writing a basic HTTP server is not too bad with the net/http package

import (
    "fmt"
    "net/http"
)

// a fundamental concept in net/http servers is 'handlers' A handler is an 
// object implementing the http.Handler interface. A common way to write a 
// handler is by using the http.HandlerFcun adapter on functions with the
// appropriate signature
func hello(w http.ResponseWriter, req *http.Request) {

    // Functions serving as handlers take a http.ResponseWriter and a 
    // http.Request as arguments. the response writer is used to fill in 
    // the HTTP response here our simple response is just "hello\n"
    fmt.Fprintf(w, "hello\n")

}

func headers(w http.ResponseWriter, req *http.Request) {

    // This handler does something a little more sophisticated by reading
    // all the HTTP request headers and echoing them into the response body
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
    // we register our handlers on server routes using the http.HandlerFunc
    // convenience function. It sets up the default router in the net/http
    // package and takes a function as an argument
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)
    
    // Finally call the ListenAndServe with the port and a handler. nil tells
    // it to use the default router we've just set up.
    http.ListenAndServe(":8090", nil)
}

// run the server in the background, and then in another terminal 
// run 'curl localhost:8090/hello'
