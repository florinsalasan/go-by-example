package main

// Go stdlib has tools for outputting logs from Go programs, with the log
// package for free-form output and the log/slog package for structured output
import (
    "bytes"
    "fmt"
    "log"
    "os"

    "log/slog"
)

func main() {

    // Invoking functions like println() from the log package uses the std
    // logger, which is already pre-configured for the reasonable logging
    // output to os.Stderr, additional methods like Fatal* or Panic* will
    // exit the program after logging, with some exceptions.
    log.Println("standard logger")

    // Loggers can be configured with flags to set their output format, by 
    // default the std logger has the log.Ldate and log.Ltime flags set, and
    // these are collected in log.LstdFlags. Can change its flags to emit
    // time with microsecond accuracy, as shown:
    log.SetFlags(log.LstdFlags | log.Lmicroseconds)
    log.Println("with micro")

    // It also supports emitting the file name and line from which the 
    // log function is called.
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Println("with file/line")
    
    // It may be useful to create a custom logger and pass it around. When 
    // creating a new logger, we can set a prefix to distinguish its 
    // output from other loggers.
    mylog := log.New(os.Stdout, "my:", log.LstdFlags)
    mylog.Println("from mylog")

    // can seet the prefix on existing loggers (including the standard one)
    // using the SetPrefix method
    mylog.SetPrefix("ohmy:")
    mylog.Println("from mylog")

    // Loggers can have custom output targets; any io.Writer works
    var buf bytes.Buffer
    buflog := log.New(&buf, "buf:", log.LstdFlags)

    // This call writes the log output into buf
    buflog.Println("hello")

    // This will actually show it on std output
    fmt.Print("from buflog:", buf.String())

    // The slog package provides structured log output. For example, logging
    // in JSON format is straightforward.
    jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
    myslog := slog.New(jsonHandler)
    myslog.Info("hi there")

    // In addition to the message, slog output can contain an arbitrary
    // number of key=value pairs
    myslog.Info("hello again", "key", "val", "age", 25)
}
// Go run the file for some sample output
