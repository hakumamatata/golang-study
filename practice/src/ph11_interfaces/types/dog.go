package types

import (
	"fmt"
	"strconv"
)

type Dog struct {
	Name string
	Age  int
}

func (d *Dog) Fly() {
	fmt.Println(d.Name, " is flying")
}

func (d *Dog) Bark(decibel int) string {
	//fmt.Println(d.Name, "Bark!!!")
	return strconv.Itoa(decibel) + "! Sound!!"
}

func (d *Dog) Run() {
	fmt.Println("Dog ", d.Name, " is Run!!!")
}
