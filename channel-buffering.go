package main

import "fmt"

// By default channels are unbuffered, meaning that they only accept sends
// (chan <-) if there is a corresponding receive (<- chan) ready to receive
// the sent value. Buffered channels accept a limited number of values 
// without a corresponding receiver for those values
func main() {

    messages := make(chan string, 2)
    // here is a channel of strings buffereing up to 2 values

    messages <- "buffered"
    messages <- "channel"
    // since the channel was buffered we can send multiple values into 
    // the channel without receiving the first one

    fmt.Println(<-messages)
    fmt.Println(<-messages)
    // we can then access the multiple values sent into the channel after
    // from what I can see the order the values are sent in are also the 
    // first values sent out, FIFO / queue

}
