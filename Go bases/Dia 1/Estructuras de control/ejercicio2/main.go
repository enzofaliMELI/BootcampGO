package main

import (
	"fmt"
)

func main() {
	var edad int
	fmt.Print("Edad: ")
	fmt.Scan(&edad)

	var empleado bool
	fmt.Print("Empleado: ")
	fmt.Scan(&empleado)

	var antig float64
	fmt.Print("Antiguedad: ")
	fmt.Scan(&antig)

	if edad > 22 && empleado && antig >= 1 {
		fmt.Println("Apto para prestamo")

		var ing int
		fmt.Println("Ingresos: ")
		fmt.Scan(&ing)
		if ing > 100000 {
			fmt.Println("No corresponde interes")
		} else {
			fmt.Println("Corresponde interes")
		}

	} else {
		fmt.Println("No apto para prestamo")
	}
}
