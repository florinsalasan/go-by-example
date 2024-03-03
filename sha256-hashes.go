package main

import (
    "crypto/sha256"
    "fmt"
)

// sha256 are used to compute short identities for binary or text blobs
// for example TLS/SSL certs use sha256 to compute a certs signature.
// heres how to compute the hashes in go

// go has multiple hash functions in various crypto/* packages

func main() {

    s := "sha256 this string"

    // start with a new hash
    h := sha256.New()

    // write expects bytes so convert the string into []byte
    h.Write([]byte(s))

    // this gets the final hash result as a byte slice, the argument to
    // Sum can be used to append to an existing byte slice, usually
    // not needed though
    bs := h.Sum(nil)

    fmt.Println(s)
    fmt.Printf("%x\n", bs)

}
// running this will compute the hash and then print it in hex, can compute
// other hashes such as SHA512 by importing it, check out the crypto go 
// package and when implementing cryptographically secure hashes, double
// check the pros and cons of different types.
