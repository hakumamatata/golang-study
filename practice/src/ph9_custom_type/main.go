package main

import (
	"fmt"
	"ph9_custom_type/types"
)

func main() {
	// pointer
	myTypeA := types.MyTypeA("Test Pointer")
	myTypeA.MethodA()
	myTypeA.MethodB() // 這邊有使用指標修改 才會影響到原本的值 而不是副本
	myTypeA.MethodA()

	// test params
	myTypeB := types.MyTypeB("MyTypeB")
	myTypeC := types.MyTypeC("MyTypeC")
	fmt.Println(myTypeB.Add(string(myTypeC)))
}
