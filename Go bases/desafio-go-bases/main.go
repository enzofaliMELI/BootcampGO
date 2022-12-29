package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("Se detectaron varios errores en tiempo de ejecución.")
		}
		fmt.Println("Ejecución finalizada.")
	}()

	filename := "tickets.csv"

	err := tickets.OpenTickets(filename)
	if err != nil {
		fmt.Println("Error in OpenTickets().")
		panic(err)
	}

	/*
		err = tickets.PrintTickets()
		if err != nil {
			fmt.Println("Error in PrintTickets().")
			panic(err)
		}
	*/

	// ejemplo 1
	pais := "Argentina"
	total, err := tickets.GetTotalTickets(pais)
	if err != nil {
		fmt.Println("Error in GetTotalTickets().")
		panic(err)
	}
	fmt.Printf("El total de personas que viajaron a %s fue de: %v \n", pais, total)

	// ejemplo 2
	tiempo := ""
	tiempo = "madrugada" // 0 → 6
	//tiempo = "mañana"    // 7 → 12
	//tiempo = "tarde" // 13 → 19
	//tiempo = "noche" // 20 → 23
	subtotal, err := tickets.GetCountByPeriod(tiempo)
	if err != nil {
		fmt.Println("Error in GetCountByPeriod().")
		panic(err)
	}
	fmt.Printf("El total de personas que viajaron de %s fue de: %v \n", tiempo, subtotal)

	// ejemplo 3
	average, err := tickets.AverageDestination(pais)
	if err != nil {
		fmt.Println("Error in AverageDestination().")
		panic(err)
	}
	fmt.Printf("El promedio de personas que viajaron a %s fue de: %v \n\n", pais, average)

}
