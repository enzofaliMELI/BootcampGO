package main

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

type CustomError struct {
	mensaje string
	code    int
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Ha ocurrido un error! %v, Codigo: %v", e.mensaje, e.code)
}

func validarSalario(salario int) (msg string, err error) {
	if salario < 10000 {
		err = CustomError{"Error: el salario es menor a 10.000", 101}
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
	if errors.Is(err, CustomError{}) {
		fmt.Println(err)
		return
	}
	fmt.Println(msg, err)

}
