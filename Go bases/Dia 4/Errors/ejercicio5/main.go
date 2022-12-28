package main

import (
	"errors"
	"fmt"
)

func calcularSalario(horas int, valor int) (salario int, err error) {

	if horas < 80 || horas < 0 {
		err = errors.New("Error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
		return
	}

	salario = horas * valor
	if salario >= 150000 {
		salario -= int(float64(salario) * 0.1)
	}
	return
}

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
	var horas int
	fmt.Print("Ingrese el horas: ")
	fmt.Scan(&horas)

	var valor int
	fmt.Print("Ingrese el valor: ")
	fmt.Scan(&valor)

	salario, err := calcularSalario(horas, valor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("El salario mensual es de %d\n", salario)

	msg, err := validarSalario(salario)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msg, err)
}
