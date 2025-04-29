package main

import (
	"fmt"
	"math/rand"
)

// func add(x int, y int) int {
func add(x, y int) int { // shortened definition because they are both the same type
	return x + y
}

func hello() {
	fmt.Println("Hello, world!")
	// print and use rand
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println(add(5, 25))

	name := "Brandon" // declare and initialize w/ string
	fmt.Println(name)

	i := 0
	for i < 28 { // basic for (i.e. while)
		i++
	}
	fmt.Println(i)

	for j := 0; j < 5; j++ {
		// do a thing
	}

	for i := range 5 { // just like range in python
		if i%2 == 0 {
			fmt.Println(i)
		} else {
			fmt.Println("Odd")
		}
	}

	for {
		// loop until hits a break
		break
	}

	if num := 9 % 2; num != 0 { // assign variable and use in condition, all in one line
		fmt.Println("ODD")
	}

	// switch demo
	switch name {
	case "Brandon":
		fmt.Printf("Hello %s!!\n", name)
	case "Sara":
		fmt.Println("Hey", name)
	default:
		fmt.Println("Hello World!")
	}
}
