package main

import (
    "fmt"
    "time"
)
// Go has a lot of support for times and durations.

func main() {

    p := fmt.Println

    // start by getting the current time
    now := time.Now()
    p(now)

    // can build a time struct by providing the year, month, day, etc. Times
    // are associated with a Location ie time zone.
    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    // can then grab the components from time individually afterwards
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    // monday-sunday weekday is also available
    p(then.Weekday())

    // the following can compare two times, testing if the first is blank
    // comapred to the second value, just read it normally
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    // the sub methods returns a duration representing the intervale between
    // two times
    diff := now.Sub(then)
    p(diff)
    // can get the diff in different units
    p(diff.Hours()/24/365.25)
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    // can also use add to advance a time by a given amount or with a '-' to 
    // move backwards a given duration
    p(then.Add(diff))
    p(then.Add(-diff))

}
