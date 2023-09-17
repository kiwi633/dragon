package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Dragon programming language!\n")
	fmt.Printf("Author: suntong\n")
	fmt.Printf("This is an interpreter program developed using Go \n")
	fmt.Printf("Go version: %s", runtime.Version())
	fmt.Printf("Feel free to type in commands\n")
}
