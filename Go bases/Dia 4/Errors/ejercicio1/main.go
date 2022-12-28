package main

import (
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
	if salario < 150000 {
		err = CustomError{"Error: el salario ingresado no alcanza el mínimo imponible", 100}
		return
	} else {
		msg = "Debe pagar impuesto"
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
