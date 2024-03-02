package main

import (
    "os"
    "text/template"
)

// Go offers builtin support for creating dynamic content or showing
// custom output to the user with text/template package. There is a sibling
// package called html/template that has the same API but with extra security
// features and should be used for generating HTML

func main() {

    // can create a new template and parse its body from a string. Templates
    // are a mix of static text and 'actions' enclosed in {{...}} that are used
    // to dynamically insert content. Similar to django templating iirc

    t1 := template.New("t1")
    t1, err := t1.Parse("value is {{.}}\n")
    if err != nil {
        panic(err)
    }

    // alternatively, can use the template.Must function to panic
    // in case Parse returns an error. This is useful for templates that 
    // were initialized in the global scope
    t1 = template.Must(t1.Parse("Value: {{.}}\n"))

    // by 'executing' the template we generate its text with specific
    // values for its actions. the {{.}} action is replaced by the value 
    // passed as a parameter to Execute
    t1.Execute(os.Stdout, "some text")
    t1.Execute(os.Stdout, 5)
    t1.Execute(os.Stdout, []string{
        "Go",
        "Rust",
        "C++",
        "C",
    })

    // A helper function for later.
    Create := func(name, t string) *template.Template {
        return template.Must(template.New(name).Parse(t))
    }

    // if the data is a struct we can use the {{.FieldName}} action to access
    // its fields. the fields should be exported to be accessible when a 
    // template is executing.
    t2 := Create("t2", "Name: {{.Name}}\n")

    t2.Execute(os.Stdout, struct {
        Name string
    }{"Jane Doe"})

    // same applies to maps; but with maps there is no restriction on the 
    // case of key names
    t2.Execute(os.Stdout, map[string]string{
        "Name": "Mickey Mouse", 
    })

    // can use if/else to provide conditional execution for templates. A
    // value is considered false if it's the default value of a type, such 
    // as 0, an empty string, nil pointer, etc. 'falsy' values just like js
    // This sample will show the usage of '-' in actions to trim whitespace

    t3 := Create("t3",
        "{{if . -}} yes {{else -}} no {{end}}\n")
    t3.Execute(os.Stdout, "not empty")
    t3.Execute(os.Stdout, "")

    // range blocks are used to loop through slices, arrays, maps or channels
    // can use {{.}} which inside of a range block is set to the current item
    // of the iteration.

    t4 := Create("t4",
        "Range: {{range .}}{{.}} {{end}}\n")
    t4.Execute(os.Stdout,
        []string{
            "Go",    
            "Rust",    
            "C++",    
            "C",    
        })


}
