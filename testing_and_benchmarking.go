package main

// Unit testing is an import part of writing principled Go programs, the 
// testing package provides the tools needed to write unit tests and the go
// test command runs tests

// for the sake of demonstration this code is in package main, could be 
// placed in any package, testing code typically lives in the same package
// as the code it tests.

import (
    "fmt"
    "testing" 
)

// will be using this simple implementation of int minimum for testing.
// Usually the code being tested would be in a source file like intutils.go
// and then the test file would be intutils_test.go
func IntMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// test is created by writing a function with a name beginning with Test
func TestIntMinBasic(t *testing.T) {
    ans := IntMin(2, -2)
    if ans != -2 {
        // t.Error* will report test failures but continues to execute the 
        // test. t.Fatal* will report failures and stop the test immediately
        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
    }
}

// writing tests can be quite repetetive, so it's idiomatic to use a 
// table-driven style, where test inputs and expected outputs are listed
// in a table and a single loop walks over them and performs the test logic
func TestIntMinTableDriven(t *testing.T) {

    var tests = []struct {
        a, b int
        want int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }

    // t.Run enables running 'subtests', one for each table entry
    // these are shown separately when executing go test -v
    for _, tt := range tests {

        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := IntMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

func BenchmarkIntMin(b *testing.B) {
    // Benchmark tests typeically go in _test.go files and are named 
    // beginning with Benchmark, the testing runner executes each benchmark
    // function several times, increasing b.N on each run until it collects
    // a precise measurement.
    // Typically the benchmark runs a function we're benchmarking in a loop
    // b.N times

    for i := 0; i < b.N; i++ {
        IntMin(1, 2)
    }

}

// run all tests with 'go test -v'

// run all benchmarks in the current project, all tests are run prior to the 
// benchmarks, the bench flag filters benchmark func names with a regexp
