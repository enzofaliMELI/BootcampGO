package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Customer struct {
	ID      int
	Name    string
	DNI     string
	Phone   string
	Address string
}

var customers []*Customer

func RegisterCustomer(id int, name string, dni string, phone string, address string) *Customer {
	customer := &Customer{
		ID:      id,
		Name:    name,
		DNI:     dni,
		Phone:   phone,
		Address: address,
	}
	return customer
}

func CheckCustomerExists(customer *Customer) bool {
	for _, c := range customers {
		if c.ID == customer.ID {
			return true
		}
	}
	return false
}

func ValidateCustomer(customer *Customer) (bool, error) {
	if customer.ID == 0 {
		return false, errors.New("Error: ID must not be zero.")
	}
	if customer.Name == "" {
		return false, errors.New("Error: Name must not be empty.")
	}
	if customer.DNI == "" {
		return false, errors.New("Error: DNI must not be empty.")
	}
	if customer.Phone == "" {
		return false, errors.New("Error: Phone must not be empty.")
	}
	if customer.Address == "" {
		return false, errors.New("Error: Address must not be empty.")
	}
	return true, nil
}

func PrintCustomers(customer []*Customer) {
	fmt.Println("Customers:")
	for _, c := range customers {
		fmt.Printf("ID: %d, Name: %s, DNI: %s, Phone: %s, Address: %s\n", c.ID, c.Name, c.DNI, c.Phone, c.Address)
	}
	fmt.Println()
}

func OpenCustomers(filename string) (err error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records:", err)
		return
	}

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		customer := &Customer{
			ID:      id,
			Name:    record[1],
			DNI:     record[2],
			Phone:   record[3],
			Address: record[4],
		}
		customers = append(customers, customer)
	}
	return
}

func SaveCustomers(customers []*Customer) (err error) {
	file, err := os.Create("customers.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}

	defer file.Close()

	for _, customer := range customers {
		_, err := fmt.Fprintf(file, "%d,%s,%s,%s,%s\n", customer.ID, customer.Name, customer.DNI, customer.Phone, customer.Address)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("Se detectaron varios errores en tiempo de ejecución.")
		}
		fmt.Println("Ejecución finalizada.")
	}()

	filename := "customers.csv"

	// Open Customers
	OpenCustomers(filename)

	// Print Customers
	PrintCustomers(customers)

	// Create a new Customers
	newCustomer := RegisterCustomer(120, "John Smith", "12345678", "12345678", "123 Main St")

	// Verificate if the Customer already exist
	if CheckCustomerExists(newCustomer) {
		panic("Error: el cliente ya existe.")
	}

	// Validate Customer
	valid, err := ValidateCustomer(newCustomer)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Customer data is valid. %v\n", valid)
		fmt.Println()
	}

	// Append Customer
	customers = append(customers, newCustomer)

	// Print Customers
	PrintCustomers(customers)

	// Save Customers
	SaveCustomers(customers)
	fmt.Println("Customers saved successfully.")
}
