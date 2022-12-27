package main

import "fmt"

func main() {

	var (
		temperatura int
		humedad     float32
		presion     float32
	)

	temperatura = 25
	humedad = 78.9
	presion = 1

	fmt.Printf("Temperatura: %v \nHumedad: %v \nPresion: %v\n", temperatura, humedad, presion)

}
