package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	reverse.String("H")
	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}

func sum(times int) {
	fmt.Println("sum")
}
