package main

import (
    "fmt"
    "math"
)

// interfaces are named collections of method signatures, similar to 
// class inheritance in other languages. This is go's way of oop like 
// implementations. Where an interface is 'implemented' by any struct
// that has methods that match the collection in a given interface
type geometry interface {
    area() float64
    perim() float64
}

// for an example before we get to the actual example in gobyexample,
// imagine a struct for dogs, and one for cats, along with an interface
// of Pet. We want to have each pet have a method that speaks, cat will 
// print out 'meow', dog will have 'woof' and we initially try making a list
// of just cat, but would not be able to append a dog struct to the list, 
// an interface lets us do that by having []Pet{} and then we can append both
// a cat and a dog to it, letting us to then loop over the slice and call speak
// on anything that is within the slice because the pet interface would 
// have been implemented and thus we know the method exists for anything inside

type rect2 struct {
    width, height float64
}

type circle struct {
    radius float64
}

// implementing the interface on rect2
func (r rect2) area() float64 {
    return r.width * r.height
}

func (r rect2) perim() float64 {
    return 2*r.width + 2*r.height
}

// implementing the interface on circle
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    // if a variable has an interface type, we can call methods that are in
    // the named interface 
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect2{width: 3, height: 4}
    c := circle{radius: 5}

    // measure can take in both rects and circles since they both implemented
    // the geometry interface
    measure(r)
    measure(c)
}
