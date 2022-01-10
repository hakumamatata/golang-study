package types

import (
	"fmt"
	"strconv"
)

type Cat struct {
	Name string
	Age  int
}

func (c Cat) Fly() {
	fmt.Println(c.Name, " is flying (Cat)")
}

func (c Cat) Bark(decibel int) string {
	return strconv.Itoa(decibel) + "! Sound!! (Cat)"
}

func (c Cat) Jump() {
	fmt.Println("Cat ", c.Name, " is Jump!!")
}
