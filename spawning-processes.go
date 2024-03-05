package main

// sometimes we need to spawn other non-go processes
import (
    "fmt"
    "io"
    "os/exec"
)

func main() {

    // start with a simple command that takes no arguments or inputs and 
    // just prints something to stdout. The exec.Command helper creates
    // an object to represent this external process
    dateCmd := exec.Command("date")

    // The Output method runs the command, waits for it to finish and collects
    // its standard output. If there were no errors, dateOut will hold bytes
    // with the date info
    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))

    // Output and other methods of Command will  return *exec.Error if there
    // was a problem executing the command (eg wrong path) and *exec.ExitError
    // if the command ran but exited with a non-zero return code
    _, err = exec.Command("date", "-x").Output()
    if err != nil {
        switch e := err.(type) {
        case *exec.Error:
            fmt.Println("failed executing:", err)
        case *exec.ExitError:
            fmt.Println("command exit rc =", e.ExitCode())
        default:
            panic(err)
        }
    }

    // Next will look at a slightly more involved case where we pipe data
    // to the external process on its stdin and collect the results from
    // its stdout
    grepCmd := exec.Command("grep", "hello")

    // Here we explicitly grab input/output pipes, start the process, write
    // some input to it, read the resulting output, and finally wait for 
    // the process to exit 
    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := io.ReadAll(grepOut)
    grepCmd.Wait()

    // Omitted error checks in this example but could use the typical 
    // if err != nil check for the different steps, also only collect the 
    // stdoutpipe results but could also collect the stderrpipe in the same
    // way
    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

    // Note that when spawining commands need to provide an explicitly 
    // delineated command and argument array, vs being able to just pass in one
    // CLI string. If wanting to spawn a full  command with a string, can
    // use bash's '-c' option;
    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
    
}
// The spawned programs return output that is the same as if we had ran
// them directly from the CLI

// date doesn't have a '-x' flag so it will exit with an error message
// and non zero return code.
