package main

import (
	"fmt"
)

func main() {
	fmt.Printf("hello %v\n", Add(1, 2))
}

func Add(a, b int) int {
	return a + b
}
