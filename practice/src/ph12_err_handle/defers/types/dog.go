package types

import "fmt"

type Dog string

func (d Dog) Run() {
	fmt.Println("Dog Run!!")
}

func (d Dog) Jump() {
	fmt.Println("Dog Jump!!!")
}
