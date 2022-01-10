package main

import (
	"fmt"
	"ph11_interfaces/interfaces"
	"ph11_interfaces/types"
)

func animalActions(a interfaces.Animal) {
	a.Fly()
	fmt.Println(a.Bark(10))

	// 型別斷言(可以調配一些行為)
	// 根據實作時參數的設置 決定是否為pointer 要加上*
	dog, ok := a.(*types.Dog)
	//fmt.Println("Is dog??", dog, ok)
	if ok {
		dog.Run()
	}
	cat, ok := a.(types.Cat)
	if ok {
		cat.Jump()
	}
}

func main() {
	dogChris := types.Dog{Name: "Chris", Age: 7}
	catRyan := types.Cat{Name: "Ryan", Age: 3}

	// 根據實作時參數的設置 決定是否為pointer 要加上&
	animalActions(&dogChris)
	fmt.Println("")
	animalActions(catRyan)
}
