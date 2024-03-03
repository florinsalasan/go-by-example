package main

import (
    // import the encoding package with the alias b64 to save a bit of space
    b64 "encoding/base64"
    "fmt"
)
// go has support for base64 encoding and decoding

func main() {

    data := "abc123!?$*&()'-=@~"
    // this is the string that will be encoded and decoded

    // Go supports standard and url compatible base64, here is how to use 
    // them. First the standard encoder which requuires a []byte rather
    // than a string, so convert data first.
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    // decoding can return errors, so check for any if it is unknown if 
    // the input is well-formed
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    // now for the url compatible encoding/decoding
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))

    // the string encodes to slightly different values between the std and url
    // encoders, but they will both decode to the original string as expected
}
