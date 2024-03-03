package main

import (
    "encoding/xml"
    "fmt"
)

// go supports xml and xml-like formats with encoding.xml package

type Plant struct {

    XMLName xml.Name `xml:"plant"`
    Id int `xml:"id,attr"`
    Name string `xml:"name"`
    Origin []string `xml:"origin"`

}
// plant will be mapped to xml, similar to the json encoding field tags
// contain directives for the encoder and decoder. here we use some special
// features of the xml package, namely the XMLName field name which 
// dictates the name of the XML element representing this struct,
// id,attr means that the Id field is an XML attribute instead of a 
// nested element

func (p Plant) String() string {

    return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
        p.Id, p.Name, p.Origin)

}

func main() {

    coffee := &Plant{Id: 27, Name: "Coffee"}
    coffee.Origin = []string{"Ethiopia", "Brazil"}

    // emit xml representing the plant using marshalindent to make it 
    // more readable.
    out, _ := xml.MarshalIndent(coffee, " ", "  ")
    fmt.Println(string(out))

    // add a generic xml header to the output by appending it explicitly
    fmt.Println(xml.Header + string(out))

    // use unmarshal to parse a stream of bytes with xml into a data structure
    // if the xml is malformed or cannot be mapped onto plant, an error
    // will be returned
    var p Plant
    if err := xml.Unmarshal(out, &p); err != nil {
        panic(err)
    }
    fmt.Println(p)

    tomato := &Plant{Id: 81, Name: "Tomato"}
    tomato.Origin = []string{"Mexico", "California"}

    // the parent>child>plant field tag tells the encoder to nest all
    // plants under parent child
    type Nesting struct {

        XMLName xml.Name `xml:"nesting"`
        Plants []*Plant `xml:"parent>child>plant"`

    }

    nesting := &Nesting{}
    nesting.Plants = []*Plant{coffee, tomato}

    out, _ = xml.MarshalIndent(nesting, " ", "  ")
    fmt.Println(string(out))

}
