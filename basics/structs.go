package main

import "fmt"

type Pizza struct {
	size     string
	toppings []string
	Style    // embedded interface
}

type Style interface {
	prepare()
}

type DeepDish struct {
	prepTime int
	cookTime int
	layers   int
}

type ThinCrust struct {
	prepTime int
	cookTime int
}

func (pizza DeepDish) prepare() {
	totalTime := pizza.prepTime + pizza.cookTime
	fmt.Printf("Your pizza will be complete in %d minutes! It has %d layers!\n",
		totalTime,
		pizza.layers,
	)
}

func (pizza ThinCrust) prepare() {
	totalTime := pizza.prepTime + pizza.cookTime
	fmt.Printf("Your pizza will be complete in %d minutes!\n", totalTime)
}

func pizzaDemo() {
	pizza1 := Pizza{
		size: "medium", toppings: []string{"Pineapple", "Ham"}, Style: DeepDish{
			prepTime: 10,
			cookTime: 25,
			layers:   5,
		},
	}
	pizza2 := Pizza{
		size: "large", toppings: []string{"Pepperoni", "Extra Cheese"}, Style: ThinCrust{
			prepTime: 4,
			cookTime: 13,
		},
	}

	pizza1.prepare()
	pizza2.prepare()
}
