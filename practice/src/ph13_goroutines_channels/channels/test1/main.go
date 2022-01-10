package main

import "fmt"

func xxx(c chan string) {
	for i := 0; i < 100; i++ {
		fmt.Print("x")
	}
	c <- "xxx"
}

func yyy(c chan string) {
	for i := 0; i < 100; i++ {
		fmt.Print("y")
	}
	c <- "yyy"
}

func main() {
	var myChannelA chan string
	myChannelA = make(chan string)
	fmt.Println(myChannelA)
	myChannelB := make(chan string)

	go xxx(myChannelA)
	go yyy(myChannelB)
	resultA := <-myChannelA
	resultB := <-myChannelB
	fmt.Println("")
	fmt.Println(resultA)
	fmt.Println(resultB)

	// 測試多次使用同一個CHAN 可以共用 並且照著安排順序執行
	go xxx(myChannelA)
	go xxx(myChannelA)
	resultC := <-myChannelA
	resultC = <-myChannelA
	fmt.Println("")
	fmt.Println(resultC)

}
