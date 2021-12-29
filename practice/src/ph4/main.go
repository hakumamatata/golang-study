package main

import (
	"fmt"
	// 自定義package
	"greeting"
)

func main() {
	fmt.Println("Hi")
	greeting.Hello()
	fmt.Println(greeting.WeekDay)
}
