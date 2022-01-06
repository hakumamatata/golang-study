package main

import "fmt"

func main() {
	/*
		output:
		1
		12
		123
		12
		1
	*/
	maxNum := 3
	ttlLen := (maxNum * 2) - 1

	for i := 1; i <= ttlLen; i++ {
		if i <= maxNum {
			for j := 1; j <= i; j++ {
				fmt.Print(j)
			}
		} else {
			for j := 1; j <= ttlLen-i+1; j++ {
				fmt.Print(j)
			}
		}
		fmt.Println("")
	}
}
