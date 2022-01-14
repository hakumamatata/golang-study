package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func sayWoo() {
	fmt.Println("Woo")
}

func firstClass(theFunc func()) {
	theFunc()
	theFunc()
}

func mathFunc(a int, b int) int {
	return a + b
}

func main() {
	firstClass(sayHi)
	firstClass(sayWoo)

	// 用宣告變數的方式 函式的參數和回傳值都需要宣告
	var myFunc func(int, int) int
	myFunc = mathFunc
	myFunc2 := mathFunc // 也可以使用短宣告
	fmt.Println(myFunc(11, 12))
	fmt.Println(myFunc2(21, 22))
}
