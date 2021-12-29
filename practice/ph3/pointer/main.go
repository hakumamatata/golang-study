package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 一般宣告
	var myInt int = 4
	var myIntPointer *int = &myInt
	fmt.Println(&myInt)
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)
	fmt.Println(reflect.TypeOf(myIntPointer))

	// 短宣告
	myFloat := 1.57
	myFloatPointer := &myFloat
	fmt.Println(&myFloat)
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)
	fmt.Println(reflect.TypeOf(myFloatPointer))

	// 替換值
	fmt.Println("Before", myInt)
	*myIntPointer = 99
	fmt.Println("After", myInt)

	// 回傳指標
	myIntAPointer := createPointer()
	fmt.Println(myIntAPointer)
	fmt.Println(*myIntAPointer)

	// 參數使用指標
	var myIntB int = 50
	printPointer(&myIntB)
}

func createPointer() *int {
	var myIntA = 9
	return &myIntA
}

func printPointer(myIntBPointer *int) {
	fmt.Println(myIntBPointer)
	fmt.Println(*myIntBPointer)
}
