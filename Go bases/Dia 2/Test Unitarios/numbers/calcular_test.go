package numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
	// Input/Output
	num1, num2 := 2, 2
	esperado := 4

	// Ejecucion
	resultado := Sumar(num1, num2)

	// Validacion

	/*
		if resultado != float64(esperado) {
			t.Fatalf("Se obtuvo %v, se esperaba %v", resultado, esperado)
		}
	*/

	assert.Equal(t, esperado, resultado, "Deben de ser iguales")
}
