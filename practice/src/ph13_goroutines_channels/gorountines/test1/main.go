package main

import (
	"fmt"
	"time"
)

var pTime int = 100

func a() {
	for i := 0; i <= pTime; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i <= pTime; i++ {
		fmt.Print("b")
	}
}

func main() {
	go a()
	go b()
	time.Sleep(time.Second)
	fmt.Println("")
	fmt.Println("main end")
}
