package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type datos struct {
	Nombre   string
	Apellido string
}

func main() {
	// Crear un router con gin
	router := gin.Default()

	// Ejercicio 1
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	// Ejercicio 2
	router.POST("/saludo", func(ctx *gin.Context) {

		var d datos

		if err := ctx.ShouldBindJSON(&d); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar los datos de entrada"})
			return
		}

		mensaje := fmt.Sprintf("Hola %s %s", d.Nombre, d.Apellido)
		ctx.JSON(http.StatusOK, gin.H{"mensaje": mensaje})
	})

	// Corremos nuestro servidor en el puerto 8080
	router.Run()
}
