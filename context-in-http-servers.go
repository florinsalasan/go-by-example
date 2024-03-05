package main

// In the last example we looked at setting up a simple HTTP server, they are
// useful for demonstrating the use of context.Context for controlling 
// cancellation a Context carries deadlines, cancellation signals, and
// other request-scoped values across API boundaries and goroutines
import (
    "fmt"
    "net/http"
    "time"
)

func hello(w http.ResponseWriter, req *http.Request) {

    // a context.Context is created for each request by the net/http package
    // and is available with the Context() method
    ctx := req.Context()
    fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")

    // wait a few seconds before sending a reply to the client. This could
    // simulate some work the server is doing. While working, keep watching
    // the Context's Done() channel for a signal that we should cancel the 
    // work and return as soon as possible.
    select {
    case <-time.After(10 * time.Second):
        fmt.Fprintf(w, "hello\n")
    case <-ctx.Done():
        // The context's Err() method returns an error that explains why
        // the Done() channel was closed
        err := ctx.Err()
        fmt.Println("server:", err)
        internalError := http.StatusInternalServerError
        http.Error(w, err.Error(), internalError)
    }
}

func main() {

    // as before register the handler on the '/hello' route, start serving
    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":8090", nil)

}
// once again run the server with go run context....go, then 
// in a different terminal window sim a client request to /hello cancelling 
// it with Ctrl+C to signal cancellation.
