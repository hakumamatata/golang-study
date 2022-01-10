package main

import (
	"errors"
	"fmt"
)

func panicHandle() {
	p := recover()
	if p != nil {
		err, ok := p.(error)
		if ok {
			fmt.Println(err.Error())
		} else {
			// 假使不是指定的型別 則可保持panic
			panic(p)
		}
	}
}

func tryOut() {
	defer panicHandle()
	func1()
}

func func1() {
	fmt.Println("func1")
	func2()
}

func func2() {
	fmt.Println("func2")
	// 如果非error的型別
	//panic("Not Error panic!!!")
	func3()
}

func func3() {
	fmt.Println("func3")
	err := errors.New("error in func3!!!")
	panic(err)
	fmt.Println("this will no show up...")
}

func main() {

	tryOut()
}
