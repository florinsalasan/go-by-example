package main

import "fmt"

// methods are defined on struct types
type rect struct {
    width, height int
}

// the area method has a *receiver type* of *rect, which is what differentiates
// methods and functions
func (r *rect) area() int {
    return r.width * r.height
}

// methods can be defined for either pointer or value receiver types.
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}
    
    // call a method with the dot operator, just like when accessing values
    // defined in a struct. similar to python again
    fmt.Println("area:", r.area())
    fmt.Println("perim:", r.perim())

    // go will also automatically handle conversion between values and pointers
    // for method calls. Kinda recommended to use a pointer receiver to avoid
    // copying on method calls or to allow the method to mutate the receiving
    // struct.
    rp := &r
    fmt.Println("area:", rp.area())
    fmt.Println("perim:", rp.perim())
}
