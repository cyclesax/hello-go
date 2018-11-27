package main

import "fmt"
import "runtime"

func main() {
    fmt.Printf("Hello, Golang World!\n")
    fmt.Printf("OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
}
