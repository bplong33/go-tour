package main

import "fmt"

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) { // multiple return vals
	return 3, 5
}

func sum(nums ...int) int { // arbitrary number of inputs
	total := 0
	for _, num := range nums { // ind, num
		total += num
	}
	return total
}

// closures!
func intSeq() func() int { // returns func that returns an int
	i := 0
	return func() int {
		i++
		return i
	}
}

// recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func funcTests() {
	// do main things

	fmt.Println(plusPlus(10, 15, 25))

	a, b := vals()
	fmt.Println("A:", a, "\nB:", b)

	s := sum(10, 115, 25)
	fmt.Println("sum: ", s)

	sequence := intSeq()
	v1 := sequence()
	fmt.Println("seq 1:", v1)
	v1 = sequence()
	fmt.Println("seq 2:", v1)
	v1 = sequence()
	fmt.Println("seq 3:", v1)

	fmt.Println("9! :", fact(9))
}
