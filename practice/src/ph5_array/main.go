package main

import "fmt"

func main() {
	// Different declarations
	var eg1 [3]int
	eg1[0] = 9
	eg1[2] = 11

	var eg2 [2]string = [2]string{"do", "re"}

	eg3 := [3]string{
		"aaaaaaaaaaaaaaaaaaaaa",
		"bbbbbbbbbbbbbbbbbbbbb",
	}

	// fmt can handle and slice, map... too
	fmt.Println(eg1)
	fmt.Printf("eg2 is %#v \n", eg2)
	fmt.Printf("eg3 is %#v \n", eg3)

	fmt.Println(len(eg1))

	// for
	for i := 0; i < len(eg3); i++ {
		fmt.Println(i)
		fmt.Println(eg3[i])
	}
	// for range
	for index, item := range eg1 {
		fmt.Println(index)
		fmt.Println(item)
	}
	for _, item := range eg2 {
		fmt.Println(item)
	}

	// 測試語法 [1:]   print => [2 3]
	testArr := []int{1, 2, 3}
	testBArr := testArr[1:]
	fmt.Println(testBArr)
	testCArr := testArr[:1]
	fmt.Println(testCArr)
}
