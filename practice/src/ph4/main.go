package main

import (
	"fmt"
	// 自定義package
	"greeting"
	// github package
	"github.com/hakumamatata/gosayhoho"
)

func main() {
	fmt.Println("Hi")
	greeting.Hello()
	fmt.Println(greeting.WeekDay)
	gosayhoho.Say()
}
