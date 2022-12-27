package main

import "fmt"

func average(numbers ...int) (average float64) {

	// Calculate the sum of the numbers
	var sum int
	for _, num := range numbers {
		sum += num
	}
	average = float64(sum) / float64(len(numbers))

	return
}

func scanInputs() (numbers []int, err bool) {

	// Scan input
	fmt.Print("Enter some number (number/n): ")
	//var numbers []int

	for {
		var n int
		_, ok := fmt.Scan(&n)
		// fmt.Println(i, ok)
		if ok != nil {
			break
		}
		if n < 0 {
			fmt.Println("El numero debe de ser positivo.")
			err = true
			break
		}
		fmt.Print("Enter other number (number/n): ")

		numbers = append(numbers, n)
	}

	return
}

func main() {
	numbers, _ := scanInputs()
	fmt.Println(average(numbers...))
}
