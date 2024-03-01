package main

import "fmt"

// Go added support for generics in v1.18, currently on 1.22 so fairly recent
// Here MapKeys takes a map of any type and returns a slice of its keys.
// This function have to type parameters, K and V. K has the comparable 
// constraint meaning the == and != operator must be useable on the keys,
// which is necessary for any key in a map in go. V has the 'any' constraint
// meaning it is not restricted in any way. any is an alias for interface{}

func MapKeys[K comparable, V any](m map[K]V) []K {
    r := make([]K, 0, len(m))
    for k := range m {
        r = append(r, k)
    }
    return r
}

// As an example of a generic type, List is a singly linked list with values
// of any type
type List[T any] struct {
    head, tail *element[T]
}

type element[T any] struct {
    next *element[T]
    val T
}

// methods can be defined on generic types just like on regular types, but 
// have to keep the type parameters in place, the type is List[T] not List
func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) GetAll() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func main() {
    var m = map[int]string{1: "2", 2: "4", 4: "8"}

    // when invoking generic functions, we can rely on type inference
    // We don't have to specify the types for K and V since the compiler
    // will infer them automatically
    fmt.Println("keys:", MapKeys(m))

    // there is nothing stopping us from specifying them explicitly if 
    // we wanted to for some reason
    _ = MapKeys[int, string](m)

    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    lst.Push(33)
    fmt.Println("list:", lst.GetAll())
}
