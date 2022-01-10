package main

import (
	"fmt"
	"ph12_err_handle/defers/types"
)

func tryOut() {
	fmt.Println("Main TryOut!!")
	dog := types.Dog("Chris")
	/*
		多個defer時 會以最後被放置的最先執行 testA() > dog.Jump() > dog.Run()
		Main testA!!!
		Dog Jump!!!
		Dog Run!!
	*/
	defer dog.Run()
	defer dog.Jump()
	defer testA()
}

func testA() {
	fmt.Println("Main testA!!!")
}

func main() {
	tryOut()
	fmt.Println("Main main!!!!!")
}
