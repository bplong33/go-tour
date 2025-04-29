package main

import "fmt"

type GasEngine struct {
	mpg           int
	gallonsInTank int
}

type ElectricEngine struct {
	mpkwh  int
	charge int
}

func (e GasEngine) milesToEmpty() int {
	return e.mpg * e.gallonsInTank
}

func (e ElectricEngine) milesToEmpty() int {
	return e.charge * e.mpkwh
}

type Engine interface {
	milesToEmpty() int
}

func engineStatus(engine Engine) {
	fmt.Println(engine)
	fmt.Println(engine.milesToEmpty())
}

func carsDemo() {
	gas := GasEngine{mpg: 25, gallonsInTank: 10}
	elec := ElectricEngine{mpkwh: 4, charge: 100}

	engineStatus(gas)
	fmt.Println("-------------------------------")
	engineStatus(elec)
}
