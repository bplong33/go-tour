package main

import (
	"fmt"
	"strings"
)

func main() {
	hello()
	fmt.Println(strings.Repeat("*", 35))

	arrayTests()
	fmt.Println(strings.Repeat("*", 35))

	mapTests()
	fmt.Println(strings.Repeat("*", 35))

	funcTests()
	fmt.Println(strings.Repeat("*", 35))

	carsDemo()
	fmt.Println(strings.Repeat("*", 35))

	pizzaDemo()
}
