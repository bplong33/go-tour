package main

import "fmt"

func rangeTest() {
	// range over built-in types

	nums := []int{2, 3, 4}
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("Total:", total)

	kv := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kv {
		fmt.Println("Key: ", k)
		fmt.Println("Value: ", v)
	}

	for i, c := range "Brandon" {
		fmt.Println(i, c)
	}
}
