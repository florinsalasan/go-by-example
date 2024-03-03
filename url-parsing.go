package main

import (
    "fmt"
    "net"
    "net/url"
)

// urls provide a uniform way to locate resources, so heres how go
// parses them

func main() {

    // we'll parse an example url, which includes a scheme, auth info, host,
    // port, path, query params, and query fragment
    s := "postgres://user:pass@host.com:5432/path?k=v#f"

    // parse and ensure no errors:
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    // accessing the scheme is fairly straight forward
    fmt.Println(u.Scheme)

    // User will contain all auth info, call Username and Password on it
    // for individual values
    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

    // the Host contains the hostname and the port, if present.
    // use SplitHostPort to extract them
    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

    // here we extract the path and the fragment after the '#'
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    // to get query params in a string of k=v format, use RawQuery, can also 
    // parse query params into a map. the parsed maps are from strings to slices
    // of strings so index into [0] if only want the first
    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])

}
