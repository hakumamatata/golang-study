package main

import (
	"fmt"
)

var (
	totalAmount float64 = 0
)

func main() {
	setLiterAmount(5.1, 6.6)
	setLiterAmount(5.1, -6.6)
	setLiterAmount(-5.1, 6.6)
	setLiterAmount(7.2, 6.675)

	fmt.Printf("All need liter is %0.2f \n", totalAmount)
}

func setLiterAmount(width float64, height float64) {
	amount, err := calculatePaintNeeded(width, height)
	if err != nil {
		// do stop
		//log.Fatal(err)
		// not stop
		fmt.Println(err)
	} else {
		fmt.Printf("when width of %0.2f and height of %0.2f \n", width, height)
		fmt.Printf("need %0.2f \n", amount)
		totalAmount += amount
	}
}

func calculatePaintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("height of %0.2f is invalid", height)
	}
	area := width * height

	return area / 10.0, nil
}
