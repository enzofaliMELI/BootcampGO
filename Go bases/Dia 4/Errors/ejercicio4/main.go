package main

import (
	"fmt"
)

func validarSalario(salario int) (msg string, err error) {
	salarioMinimo := 10000
	if salario < 10000 {
		err = fmt.Errorf("Error: el mínimo imponible es de: %d y el salario ingresado es de: %d", salarioMinimo, salario)
		return
	} else {
		msg = "OK"
		return
	}
}

func main() {
	var salario int
	fmt.Print("Ingrese el salario: ")
	fmt.Scan(&salario)

	msg, err := validarSalario(salario)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msg, err)
}
