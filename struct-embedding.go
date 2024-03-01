package main

import "fmt"

// Go supports the embedding of structs and interfaces to express a more
// seamless composition of types. This is not //go:embed which is a go
// directive to embed files and folders in the app binary

type base struct {
    num int 
}

func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

// a container embeds a base struct, embedding looks like a field
// without a name.
type container struct {
    base
    str string
}

func main() {
    // when creating structs with literals have to initialize the embedding
    // explicitly. there the embedded type is the field name.
    co := container{
        base: base {
            num: 1,
        },
        str: "some name",
    }
    // when we want to access an embedded field, can do so from the struct
    // that is embedding the other as can be seen by grabbing co.num straight
    // from the container rather than needing to access it from the embedded
    // struct.
    fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

    // however if you want to be more verbose you can access it that way too
    fmt.Println("also num:", co.base.num)

    // since container embeds base, we can call methods defined on base, on
    // container as well as embedded methods become methods of the
    // embedding struct too.
    fmt.Println("describe:", co.describe())

    type describer interface {
        describe() string
    }

    // embedding structs with methods may be used to bestow interface
    // implementations onto other structs. Here we can see how container
    // implements the describer interface because the embedded base struct
    // implemented the interface
    var d describer = co
    fmt.Println("describer:", d.describe())

}
