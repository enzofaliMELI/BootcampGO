package main

import (
	"encoding/json"
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

	router.POST("/saludito", func(c *gin.Context) {

		data, err := c.GetRawData()
		if err != nil {
			panic(err)
		}

		var d datos

		if err := json.Unmarshal(data, &d); err != nil {
			c.String(http.StatusBadRequest, "Error")
			return
		}
		mensaje := fmt.Sprintf("Hola %s %s", d.Nombre, d.Apellido)
		c.String(http.StatusOK, mensaje)
	})

	// Corremos nuestro servidor en el puerto 8080
	router.Run()
}
