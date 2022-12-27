package main

import (
	"fmt"
)

func main() {

	// Pedimos al usuario que ingrese una palabra
	var salario int
	fmt.Print("Ingrese el salario: ")
	fmt.Scan(&salario)

	// Switch statement with both
	switch {
	case salario > 150:
		fmt.Println(salario, float32(salario)*(1-0.27), 27)
	case salario > 50:
		fmt.Println(salario, float32(salario)*(1-0.17), 17)
	default:
		fmt.Println("Salario menor a $50.000")
	}

}
