package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {

    const s = "สวัสดี"
    // s is assigned a literal value representing hello in Thai. Go string
    // literals are UTF-8 encoded text

    fmt.Println("len:", len(s))
    // since strings are equivalent to []byte, this len(string) will produce 
    // the length of raw bytes stored in the string. 

    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
        // This printf will print out the hex values of all of the bytes
        // that are in the string s.
    }
    fmt.Println()

    // to instead count the number of runes, chars in other languages, use
    // the utf8 package
    fmt.Println("Rune count:", utf8.RuneCountInString(s))
    // run time of the rune count is dependent on the size of the string,
    // because it must decode each utf8 rune sequentially and some non latin
    // runes are represented by multiple utf8 code points, so the result 
    // may not be the one that is expected.

    // lets print the runes out using the range loop
    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
        // %#U is for printing unicode with character, there are reference
        // sheets for different 'verbs' to use with printf and the like
        // the range loop handles strings by decoding each rune along with
        // its offset in the string.
    }

    // Can do the same thing using a function in the utf8 package
    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width

        // can also pass a rune value into functions
        examineRune(runeValue)
    }

}

func examineRune(r rune) {
    // values enclosed in single quotes are rune literals, and can be 
    // compared to a rune value
    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}
