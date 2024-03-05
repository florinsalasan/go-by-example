package main

// previously looked at spawning external processes, done when an external
// process is needed to be accessible to a running Go process. Sometimes jsut
// want to completely replace the current Go process with another (perhaps non
// -Go) one. To do this we'll use Go's implementation of the classic exec func

import (
    "os"
    "os/exec"
    "syscall"
)

func main() {

    // for the example will exec ls, Go requires an absolute path to the 
    // binary we want to execute, so we'll use exec.LookPath to find it
    // Likely in /bin/ls
    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }

    // Exec requires arguments in slice form (as opposed to one big string)
    // will give ls a few common arguments, note that the first argument
    // should be the program name.
    args := []string{"ls", "-a", "-l", "-h"}

    // Exec also needs a set of env variables to use, we'll provide our 
    // system environment
    env := os.Environ()

    // Here is the actual syscall.Exec call, if this call is successful, the
    // execution of our process will end here and be replaced by the
    // /bin/ls -a -l -h process (assuming the absolute path is /bin/ls). If 
    // there is an error, we'll instead get a return value
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}

// When running the program, it will be replaced by ls

// Note that go does not offer a classic Unix fork function, usually not
// an issue since goroutines, spawning processes, and exec'ing processes
// covers most of the use cases for fork
