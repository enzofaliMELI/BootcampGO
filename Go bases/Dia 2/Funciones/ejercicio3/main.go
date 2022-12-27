package main

import (
	"fmt"
	"strings"
)

func scanInputs() (categoria string, horasTrabajadas int) {

	fmt.Print("Ingresar Categoria (A, B o C/n): ")

	// Validar tipo de datos
	_, err := fmt.Scan(&categoria)
	if err != nil {
		fmt.Println("Categorias deben ser A, B o C!")
		return
	}

	// Validar categoria A, B o C
	valList := []string{"A", "B", "C"}
	if !(strings.Contains(strings.Join(valList, ""), categoria)) {
		fmt.Println("Categorias deben ser A, B o C!")
		return
	}

	// Validar tipo de datos
	fmt.Print("Enter Minutos Trabajados (number/n): ")
	_, err2 := fmt.Scan(&horasTrabajadas)
	if err2 != nil {
		fmt.Println("Ingresar numero entero positivo!.")
		return
	}

	// Validar entero positivo
	if horasTrabajadas <= 0 {
		fmt.Println("El numero debe de ser entero positivo!.")
		return
	}

	return
}

func main() {
	categoria, horasTrabajadas := scanInputs()
	fmt.Printf("Categoria: %v, Horas Trabajadas: %v\n", categoria, horasTrabajadas)

	switch categoria {
	case "A":
		fmt.Println("Salario: $", float32(horasTrabajadas)*3000*1.5)
	case "B":
		fmt.Println("Salario: $", float32(horasTrabajadas)*1500*1.2)
	case "C":
		fmt.Println("Salario: $", horasTrabajadas*1000)
	}

}
