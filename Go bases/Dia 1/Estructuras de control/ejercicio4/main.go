package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println(employees["Benjamin"])

	total := 0
	for _, value := range employees {
		if value > 21 {
			total += 1
		}
	}
	fmt.Println("Total de empleados mayores de 21 años: ", total)
	fmt.Println(employees)

	employees["Federico"] = 25
	fmt.Println(employees)

	delete(employees, "Pedro")
	fmt.Println(employees)
}
