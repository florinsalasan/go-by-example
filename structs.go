package main

import "fmt"

// a struct is a collection of fields, similar to structs in c, useful for
// grouping data together 
type person struct {
    name string
    age int
}

// can safely return a pointer to a local variable since a local 
// variable will survive the scope of the function.
func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p
}

func main() {

    // this syntax creates a new struct of type person
    fmt.Println(person{"Bob", 20})

    // can name the fields when initializing the struct
    fmt.Println(person{name: "Alice", age: 40})

    // omitted fields will be zero-valued
    fmt.Println(person{name: "Fred"})

    // an & prefix yields a pointer to the struct
    fmt.Println(&person{name: "Ann", age: 50})

    // it is idiomatic to encapsulate new struct creations in a 
    // constructor function
    fmt.Println(newPerson("Jon"))

    // can access struct fields with the dot operator, ie person1.name
    s := person{name: "Sean", age: 55}
    fmt.Println(s.name)

    // can use the dot operator with struct pointers as well, and the 
    // pointers are automatically dereferenced for ease of use
    sp := &s
    fmt.Println(sp.age)

    // structs are mutable
    sp.age = 51
    fmt.Println(sp.age)

    // if a struct type will only be used for a single value, it isn't 
    // necessary to give it a name. instead it can have an anonymous struct
    // type. this is commonly used in table-driven tests which will be 
    // discussed later in this 
    dog := struct {
        name string
        isGood bool
    }{
        "Rex",
        true,
    }
    fmt.Println(dog)
}
