package main

import "fmt"

const (
	dog     = "dog"
	cat     = "cat"
	arana   = "arana"
	hamster = "hamster"
)

type Operacion func(numb float64) float64

func catFood(numb float64) (alimento float64) {
	alimento = 5 * numb
	return
}

func dogFood(numb float64) (alimento float64) {
	alimento = 10 * numb
	return
}

func aranaFood(numb float64) (alimento float64) {
	alimento = 0.25 * numb
	return
}

func hamsterFood(numb float64) (alimento float64) {
	alimento = 0.15 * numb
	return
}

func animal(operador string) (result Operacion, msg string) {
	switch operador {
	case dog:
		result = catFood
	case cat:
		result = dogFood
	case arana:
		result = aranaFood
	case hamster:
		result = hamsterFood
	default:
		msg = "Animal Invalido"
	}

	return
}

func main() {

	animalDog, msg := animal(dog)
	animalCat, msg := animal(cat)

	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)

	fmt.Println(amount, msg)
}
