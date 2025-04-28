package main

import (
	"fmt"
	"maps"
)

func mapTests() {
	m := make(map[string]int) // make map w/ string key and int val
	m["age"] = 28

	fmt.Println("map: ", m)

	v1 := m["age"]
	fmt.Println("age: ", v1)
	delete(m, "age")
	fmt.Println("map: ", m)

	m["age"] = 28
	// check if a key exists in the map
	_, v2 := m["name"]
	fmt.Println("Exists:", v2)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map n:", n)
	c := map[string]int{"foo": 1, "bar": 2}

	if maps.Equal(n, c) {
		fmt.Println("Matching!")
	}
}
