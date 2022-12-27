package main

import (
	"fmt"
)

type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (a Alumnos) Detalle() {
	fmt.Println("Nombre:", a.Nombre)
	fmt.Println("Apellido:", a.Apellido)
	fmt.Println("DNI:", a.DNI)
	fmt.Println("Fecha:", a.Fecha)
}

func main() {
	estudiante := Alumnos{
		Nombre:   "Juan",
		Apellido: "Pérez",
		DNI:      "12345678",
		Fecha:    "01/01/2020",
	}
	estudiante.Detalle()
}
