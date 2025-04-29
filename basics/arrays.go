package main

import (
	"fmt"
	"slices"
)

func arrayTests() {
	var arr [10]int
	fmt.Println("emp:", arr)

	arr[4] = 100                    // set ind 4 to 100
	fmt.Println("Index 4:", arr[4]) // get index 4

	b := [5]int{1, 2, 3, 4, 5} // declare and populate array
	fmt.Println(b)
	c := [...]int{1, 2, 3, 4, 5, 6, 7, 8} // [...] counts the number of values being passed on initialize
	fmt.Println(c)
	d := [...]int{100, 3: 400, 500} // sets ind 3 to 400, fills with 0-value
	fmt.Println("idx:", d)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	// --------- Slices --------------
	var s []string // typed by what they contain, but not pre-defined length

	s = make([]string, 5) // creates a slice w/ 5 empty indices
	fmt.Println(s, "len:", len(s), "cap:", cap(s))

	s = make([]string, 0) // creates a slice w/ 0 empty indices

	s = append(s, "Brandon", "Sara", "Oliver", "Finn", "Emmy")
	fmt.Println("s: ", s)

	cop := make([]string, len(s))
	copy(cop, s)

	l := s[2:4]
	fmt.Println("l: ", l)

	if slices.Equal(s, cop) {
		fmt.Println("Matching!")
	}

	// twoDtwo := make([][]int, 3)
	var twoDtwo = make([][]int, 3)
	for i := 0; i < len(twoDtwo); i++ {
		twoDtwo[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			twoDtwo[i][j] = i + j
		}
	}
	fmt.Println(twoDtwo)

	spreadEx := []int{1, 2, 3, 4}
	additions := []int{5, 6}
	spreadEx = append(spreadEx, additions...)
	fmt.Println("Spread Example:", spreadEx)
}
