package main

import (
	"fmt"
)

const (
	small = "pequeño"
	med   = "mediano"
	big   = "grande"
)

// Interfaz producto con funcionalidad de retornar precio
type Producto interface {
	Precio() float64
}

// Estructura del producto: Tipo y Precio
type Product struct {
	// Producto
	Type  string
	Price float64
}

// Implementacion implicita de la interfaz Producto
// Implementacion del metodo precio: Precio() float64
func (p *Product) Precio() float64 {
	var cost float64
	switch p.Type {
	case small:
		cost = p.Price
	case med:
		cost = p.Price + p.Price*0.03
	case big:
		cost = p.Price + p.Price*0.06 + 2500
	}
	return cost
}

// Factory: recibe el tipo de producto y el precio como parámetros y
// retorna una instancia de Product que implementa la interfaz Producto (metodo Precio)

func NewProduct(tipo string, precio float64) Producto {
	return &Product{Type: tipo, Price: precio}
}

func main() {

	p1 := NewProduct("pequeño", 1000)
	fmt.Println(p1.Precio()) // 1000

	p2 := NewProduct("mediano", 1000)
	fmt.Println(p2.Precio()) // 1030

	p3 := NewProduct("grande", 1000)
	fmt.Println(p3.Precio()) // 5630

}
