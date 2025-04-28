package main

import "fmt"

func main() {
	// hello()
	// arrayTests()
	// mapTests()
	// funcTests()
	rangeTest()
	name := "Brandon"
	for _, char := range name {
		fmt.Printf("Char: %c\n", char)
	}
}
