package main

import (

    "encoding/json"
    "fmt"
    "os"

)
// Go offers support for JSON encoding and decoding, including on 
// both built in and custom datatypes

// we'll use the following two structs to demonstrate encoding and 
// decoding of custom types
type response1 struct {
    Page int
    Fruits []string
}

// only exported fields will be encoded/decoded in JSON, fields must
// start with a capital letter to be exported
type response2 struct {
    Page int `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {

    // first look at encoding basic data types to JSON strings, here are some
    // examples for atomic values
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))
    
    // Some slices and maps as well:
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    // json package can also encode the custom data types we define. it will
    // only include exported fields and defaults the field names as the 
    // json keys
    res1D := &response1{
        Page: 1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    // can use tags on struct field declarations to customize the json 
    // key names as we did with the response2 struct
    res2D := &response2{
        Page: 1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))


    // after encoding let's look at how to decode it back into Go values.
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

    // need to provide a var where the json package can put the decoded data.
    // this map[string]interface{} can hold a map of strings to arbitrary
    // data types
    var dat map[string]interface{}

    // here is the decoding, along with a quick check for errors
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    // in order to use the values that are now in the decoded map, we still
    // need to convert them into their appropriate type. ie, converting the 
    // value in num to the expected float64 type
    // did not realize this was the syntax to convert values, maybe 
    // specific to convert from interface to a concrete type?
    num := dat["num"].(float64)
    fmt.Println(num)

    // getting nested data is more painful, as each level requires it's 
    // own conversion before going deeper into the nest.
    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)

    // can decode json into custom data types, this has the advantages of
    // extra type safety in a program and eliminating the need for 
    // type assertions when accessing the decoded data
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    // We've been using strings and byte slices as intermediates between
    // the data and the json representations on stdout, can stream json
    // encodings directly to os.Writers like os.Stdout or even HTTP responses
    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}
