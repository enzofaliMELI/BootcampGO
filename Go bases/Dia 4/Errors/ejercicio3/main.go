package main

import (
	"errors"
	"fmt"
)

var ErrorSalario = errors.New("Error: el salario ingresado no alcanza el mínimo imponible")

func validarSalario(salario int) (msg string, err error) {
	if salario < 10000 {
		err = ErrorSalario
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

	// La validación debe ser hecha con la función “Is()”.
	msg, err := validarSalario(salario)
	if errors.Is(err, ErrorSalario) {
		fmt.Println(err)
		return
	}
	fmt.Println(msg, err)
}
