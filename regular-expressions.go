package main

import (

    "bytes"
    "fmt"
    "regexp"

)
// Go has builtin regex support with the package regexp. Here are some 
// common regex tasks.

func main() {

    // tests whether a pattern matches a string
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

    // above the string pattern was used directly, for other regex
    // tasks will need to Compile an optimized regexp struct
    r, _ := regexp.Compile("p([a-z]+)ch")

    // lots of methods are available on the regexp structs. here is a 
    // match test like the first example we did
    fmt.Println(r.MatchString("peach"))

    // this will find the first match for the regexp
    fmt.Println(r.FindString("peach punch"))

    // this will also find the first match but returns the start
    // and end indices for the match rather than the text.
    fmt.Println("idx:", r.FindStringIndex("peach punch"))

    // submatch variants include info about both whole-pattern matches 
    // and submatches within those matches. For example will return info for 
    // both the p([a-z]+)ch and ([a-z]+) patterns
    fmt.Println(r.FindStringSubmatch("peach punch"))

    // finds indices of matches and submatches
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))

    // using the 'All' variants of the functions will apply to all matches in
    // the input rather than just the first. Obviously used when all matches
    // are wanted. The int that is passed in after the string indicates the 
    // number of matches the function should look for, if it is >= 0, will
    // return the indicated number of matches, otherwise will return all.
    fmt.Println(r.FindAllString("peach punch pinch", -1))

    // the 'All' variants are available in the other functions from earlier
    fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

    // again passing in a non negative int to an 'all' variant will limit the 
    // number of matches that the function call will look for.
    fmt.Println(r.FindAllString("peach punch pinch", 2))

    // The examples above has string args and used names like MatchString,
    // []byte args can be given as well allowing us to call similar methods
    // that do not have the 'string' keyword in the function name
    fmt.Println(r.Match([]byte("peach")))

    // when creating global vars with regexp can use MustCompile instead of
    // just Compile. This will panic in the case of an error, making
    // it safer to use for globals
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println("regexp:", r)

    // the regexp package can be used to replace subsets of strings with 
    // other values as well.
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    // the func variant allows the transformation of matched text with a 
    // given function
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}
