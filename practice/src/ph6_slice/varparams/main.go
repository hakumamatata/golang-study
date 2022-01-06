package main

import "fmt"

func getAverage(numbers ...float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func getSum(numbers ...float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	fmt.Println(getAverage(10.1, 12, 54.9))
	fmt.Println(getAverage(99.901, 45.68))

	fmt.Println(getSum(100.1, 99.54, 75.23))

	var sliceA []float64
	sliceA = append(sliceA, 98.5, 97.1)
	fmt.Println(getSum(sliceA...))
}
