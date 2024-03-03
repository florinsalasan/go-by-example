package main

import (
    "fmt"
    "time"
)

// common requirement is getting the number of seconds/milliseconds or even
// nanoseconds since the Unix epoch, this is set as 'Jan 1st 1970' UTC 00:00:00

func main() {

    // can use time.Now and then call Unix, UnixMilli, or UnixNano to 
    // get the elapsed time since Unix epoch in the respective units
    now := time.Now()
    fmt.Println(now)

    fmt.Println(now.Unix())
    fmt.Println(now.UnixMilli())
    fmt.Println(now.UnixNano())

    // can also convert integer seconds or nanoseconds since epoch into 
    // the corresponding time struct
    fmt.Println(time.Unix(now.Unix(), 0))
    fmt.Println(time.Unix(0, now.UnixNano()))
}
