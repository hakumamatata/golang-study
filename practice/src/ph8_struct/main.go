package main

import (
	"fmt"
	"ph8_struct/structs"
)

type cat struct {
	name string
	age  int
}

type dog struct {
	name string
	age  int
}

func main() {
	var structA struct {
		name string
		age  int
	}
	structA.name = "Cash"
	structA.age = 99
	fmt.Println(structA)

	catA := cat{
		name: "mary",
		age:  5,
	}
	var dogA dog
	dogA.name = "cindy"
	dogA.age = 7
	fmt.Println(catA)
	fmt.Println(dogA)

	var dogPointer *dog = &dogA
	fmt.Println(dogPointer)
	fmt.Println(dogPointer.name)
	fmt.Println((*dogPointer).name)

	var carA structs.Car
	carA.Speed = 101.78
	fmt.Println(carA)

	carB := structs.Car{Speed: 199}
	fmt.Println(carB)

	// 引用&&崁入struct使用
	// 取一個屬性名稱 引用 (好處是視覺化 以及 名稱比較不會打架)
	carA.Address.City = "Taichung"
	carA.Address.Town = "Xientun"
	// 匿名屬性 可以直接呼叫 或者是透過struct名稱 (好處是有些情況時 方便使用)
	carA.Profile.Nickname = "Mars"
	carA.Age = 5
	fmt.Println(carA)
	fmt.Printf("%#v\n", carA)

	// 測試另外一種方式 給值
	carB.Address = structs.Location{City: "YYY", Town: "UUU"}
	fmt.Println(carB.Address)
}
