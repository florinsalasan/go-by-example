package main

import "fmt"

// basic sends and receives on channels are blocking. However, we can use
// select with a default clause to implement non-blocking sends, receives,
// and even non-blocking multi-way selects

func main() {
    messages := make(chan string)
    signals := make(chan string)

    // here is a non-blocking receive, if a value is available on
    // messages then select will take the <-messages case with that value
    // otherwise it will immediately take the default case.
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    // a non blocking send works similarily, here msg cannot be sent to
    // the messages channel, because the channel has no buffer and there
    // is no receiver. thus the default case will be selected.
    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

    // we can use miltiple cases above the default clause to implement a 
    // multi-way non blocking select, here we make non-blocking receives on
    // both messages and signals
    select {
    case msg := <-messages:
        fmt.Println("sent message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}
