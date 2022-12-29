package tickets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalTickets(t *testing.T) {
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
	assert.NotNil(t, errAr)
	assert.Equal(t, esperadoAr, resultadoAr, "Deben de ser iguales")

	// Caso 2:
	assert.NotNil(t, errBr)
	assert.Equal(t, esperadoBr, resultadoBr, "Deben de ser iguales")
}

func TestGetCountByPeriod(t *testing.T) {
	// Input/Output
	// Caso 1:
	madrugada := "madrugada"
	esperadoMadrugada := 151

	// Caso 2:
	mañana := "mañana"
	esperadoMañana := 289

	// Caso 3:
	tarde := "tarde"
	esperadoTarde := 256

	// Caso 4:
	noche := "noche"
	esperadoNoche := 304

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
	assert.NotNil(t, errMadrugada)
	assert.Equal(t, esperadoMadrugada, resultadoMadrugada, "Deben de ser iguales")

	// Caso 2:
	assert.NotNil(t, errMañana)
	assert.Equal(t, esperadoMañana, resultadoMañana, "Deben de ser iguales")

	// Caso 3:
	assert.NotNil(t, errTarde)
	assert.Equal(t, esperadoTarde, resultadoTarde, "Deben de ser iguales")

	// Caso 4:
	assert.NotNil(t, errNoche)
	assert.Equal(t, esperadoNoche, resultadoNoche, "Deben de ser iguales")
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
	assert.NotNil(t, errAr)
	assert.Equal(t, esperadoAr, resultadoAr, "Deben de ser iguales")

	// Caso 2:
	assert.NotNil(t, errBr)
	assert.Equal(t, esperadoBr, resultadoBr, "Deben de ser iguales")
}
