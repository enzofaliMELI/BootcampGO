package tickets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalTickets(t *testing.T) {

	filename := "tickets.csv"
	OpenTickets(filename)

	// Input/Output
	// Caso 1:
	destinationAr := "Argentina"
	esperadoAr := 15

	// Caso 2:
	destinationBr := "Brazil"
	esperadoBr := 45

	// Ejecucion
	// Caso 1:
	resultadoAr, errAr := GetTotalTickets(destinationAr)

	// Caso 2:
	resultadoBr, errBr := GetTotalTickets(destinationBr)

	// Validacion
	// Caso 1:
	if errAr != nil {
		fmt.Println(errAr)
	} else {
		assert.Equal(t, esperadoAr, resultadoAr, "Deben de ser iguales")
	}

	// Caso 2:
	if errBr != nil {
		fmt.Println(errBr)
	} else {
		assert.Equal(t, esperadoBr, resultadoBr, "Deben de ser iguales")

	}
}

func TestGetCountByPeriod(t *testing.T) {

	// Input/Output
	// Caso 1:
	madrugada := "madrugada"
	esperadoMadrugada := 304

	// Caso 2:
	mañana := "mañana"
	esperadoMañana := 256

	// Caso 3:
	tarde := "tarde"
	esperadoTarde := 289

	// Caso 4:
	noche := "noche"
	esperadoNoche := 151

	// Ejecucion
	// Caso 1:
	resultadoMadrugada, errMadrugada := GetCountByPeriod(madrugada)

	// Caso 2:
	resultadoMañana, errMañana := GetCountByPeriod(mañana)

	// Caso 3:
	resultadoTarde, errTarde := GetCountByPeriod(tarde)

	// Caso 4:
	resultadoNoche, errNoche := GetCountByPeriod(noche)

	// Validacion
	// Caso 1:
	if errMadrugada != nil {
		fmt.Println(errMadrugada)
	} else {
		assert.Equal(t, esperadoMadrugada, resultadoMadrugada, "Deben de ser iguales")

	}

	// Caso 2:
	if errMañana != nil {
		fmt.Println(errMañana)
	} else {
		assert.Equal(t, esperadoMañana, resultadoMañana, "Deben de ser iguales")

	}

	// Caso 3:
	if errTarde != nil {
		fmt.Println(errTarde)
	} else {
		assert.Equal(t, esperadoTarde, resultadoTarde, "Deben de ser iguales")

	}

	// Caso 4:
	if errNoche != nil {
		fmt.Println(errNoche)
	} else {
		assert.Equal(t, esperadoNoche, resultadoNoche, "Deben de ser iguales")

	}
}

func TestAverageDestination(t *testing.T) {

	// Input/Output
	// Caso 1:
	destinationAr := "Argentina"
	esperadoAr := 0.015

	// Caso 2:
	destinationBr := "Brazil"
	esperadoBr := 0.045

	// Ejecucion
	// Caso 1:
	resultadoAr, errAr := AverageDestination(destinationAr)

	// Caso 2:
	resultadoBr, errBr := AverageDestination(destinationBr)

	// Validacion
	// Caso 1:
	if errAr != nil {
		fmt.Println(errAr)
	} else {
		assert.Equal(t, esperadoAr, resultadoAr, "Deben de ser iguales")

	}

	// Caso 2:
	if errBr != nil {
		fmt.Println(errBr)
	} else {
		assert.Equal(t, esperadoBr, resultadoBr, "Deben de ser iguales")

	}
}
