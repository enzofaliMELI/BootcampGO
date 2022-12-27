package main

import (
	"fmt"
	"time"
)

type Person struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}

type Employee struct {
	ID       int
	Position string
	Person
}

func (e *Employee) PrintEmployee() {
	fmt.Printf("ID: %d\n", e.ID)
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Date of birth: %s\n", e.DateOfBirth.Format("02-01-2006"))
	fmt.Printf("Position: %s\n", e.Position)
}

func main() {
	p := Person{ID: 1, Name: "John Smith", DateOfBirth: time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC)}
	e := Employee{ID: 2, Position: "Manager", Person: p}
	e.PrintEmployee()
}
