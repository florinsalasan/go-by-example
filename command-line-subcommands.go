package main

// some cli tools like the 'go' tool or 'git' have many subcommands, each with
// its own set of flags. For example 'go build' and 'go get' are two different
// subcommands of the go tool. The flag package allows us to easily define
// simple subcommands that have their own flags
import (
    "flag"
    "fmt"
    "os"
)

func main() {

    // declare a subcommand using the NewFlagSet function, and proceed to 
    // define new flags specific for this subcommand
    fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
    // both of these will be flags specific to the foo subcommand
    fooEnable := fooCmd.Bool("enable", false, "enable")
    fooName := fooCmd.String("name", "", "name")

    // for a different subcommand can define different supported flags
    barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
    // the level command will be specific to subcommand bar
    barLevel := barCmd.Int("level", 0, "level")

    // the subcommand is expected as the first argument to the program
    if len(os.Args) < 2 {
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }

    // check which subcommand is invoked
    switch os.Args[1] {

    // for each subcommand we parse its own flags and have access to 
    // trailing positional arguments.
    case "foo":
        fooCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'foo'")
        fmt.Println("  enable:", *fooEnable)
        fmt.Println("  name:", *fooName)
        fmt.Println("  tail:", fooCmd.Args())
    case "bar":
        barCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'bar'")
        fmt.Println("  level:", *barLevel)
        fmt.Println("  tail:", barCmd.Args())
    default:
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }

}

