package main

import (
	"fmt"
)

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
	Category    int
}

var Products []*Product

func (p *Product) Save(Product) {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Println(*product)
	}
}

func getByld(ID int) *Product {
	for _, prod := range Products {
		if prod.ID == ID {
			return prod
		}
	}
	panic("Articulo no econtrado")
}

func main() {

	p1 := Product{ID: 1, Name: "Lapiz", Price: 2.00, Description: "Color rojo", Category: 1}
	p1.Save(p1)
	p1.GetAll()
	p1.Name = "Aaaaaaaaaa" // p1 = Product{ID: 1, Name: "AAAAAAALapiz", Price: 2.00, Description: "Color rojo", Category: 1}
	p1.GetAll()
	fmt.Println("---------------------")

	p2 := Product{ID: 2, Name: "Lapiz", Price: 2.00, Description: "Color azul", Category: 2}
	p2.Save(p2)
	p2.GetAll()
	fmt.Println("---------------------")

	fmt.Println(*getByld(1))
	fmt.Println(*getByld(2))
	fmt.Println(*getByld(3))
}
