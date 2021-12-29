package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
	var month = now.Month()

	fmt.Println(year)
	fmt.Println(month)
}
