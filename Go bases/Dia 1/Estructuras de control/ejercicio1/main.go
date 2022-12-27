package main

import (
	"fmt"
)

func main() {
	var palabra string
	fmt.Print("Ingresa una palabra: ")
	fmt.Scan(&palabra)
	fmt.Println("Largo de la palabra", len(palabra))

	for _, letra := range palabra {
		// fmt.Println(string(letra))
		fmt.Printf("%c\n", letra)
	}
}
