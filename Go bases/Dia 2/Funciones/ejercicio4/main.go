package main

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

type Operacion func(numbers []float64) float64

func Minimum(numbers []float64) (result float64) {
	result = numbers[0]
	for _, num := range numbers {
		if num < result {
			result = num
		}
	}
	return
}

func Average(numbers []float64) (result float64) {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	result = float64(sum) / float64(len(numbers))
	return
}

func Maximum(numbers []float64) (result float64) {
	result = numbers[0]
	for _, num := range numbers {
		if num > result {
			result = num
		}
	}
	return
}

func operation(operador string) (result Operacion, err bool) {
	switch operador {
	case minimum:
		result = Minimum
	case average:
		result = Average
	case maximum:
		result = Maximum
	}

	return
}

func main() {

	minFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)

	examplelist := []float64{2, 3, 4, 5, 6}

	minValue := minFunc(examplelist)
	averageValue := averageFunc(examplelist)
	maxValue := maxFunc(examplelist)

	fmt.Println(minValue, averageValue, maxValue, err)
}
