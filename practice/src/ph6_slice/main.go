package main

import "fmt"

func main() {
	var sliceA []int
	// get err...  use append!
	//sliceA[0] = 1
	//sliceA[1] = 2

	sliceB := []int{
		2,
		4,
		6,
	}

	sliceC := make([]string, 2)
	sliceC[0] = "aa"
	sliceC[1] = "bb"

	// 如果由陣列這邊做出切片slice 則兩邊修改時 皆會影響到對方
	arrA := [3]string{"hi", "hello", "hola"}
	sliceD := arrA[0:2]

	fmt.Println(sliceA)
	fmt.Println(len(sliceA))
	fmt.Println(sliceB)
	fmt.Println(len(sliceB))
	fmt.Println(sliceC)
	fmt.Println(len(sliceC))
	fmt.Println(sliceD)
	fmt.Println(len(sliceD))
	// 如果由陣列這邊做出切片slice 則兩邊修改時 皆會影響到對方
	arrA[0] = "hi2"
	sliceD[1] = "hello2"
	fmt.Println(arrA)
	fmt.Println(sliceD)

	// 建議使用由make或者切片字面量建置 比較不會影響
	// 使用append時 盡量回傳至同一個slice變數 避免掉無法預期的結果
	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2")
	s3 := append(s2, "s3")
	s4 := append(s3, "s4")
	fmt.Println(s1, s2, s3, s4) // [s1 s1] [s1 s1 s2] [s1 s1 s2 s3] [s1 s1 s2 s3 s4]
	s3[0] = "s1Bs3"             // 在此修改會影響到append注入的s2
	fmt.Println(s1, s2, s3, s4) // [s1 s1] [s1Bs3 s1 s2] [s1Bs3 s1 s2 s3] [s1 s1 s2 s3 s4]
	s4[1] = "s1Bs4"             // 但是後續更改s4時 又沒有影響至s3......
	fmt.Println(s1, s2, s3, s4) // [s1 s1] [s1Bs3 s1 s2] [s1Bs3 s1 s2 s3] [s1 s1Bs4 s2 s3 s4]

	// 切片與零值
	var sEmpty []string
	fmt.Printf("sEmpty: %#v\n", sEmpty)
	if len(sEmpty) == 0 {
		sEmpty = append(sEmpty, "first item")
	}
	fmt.Printf("sEmpty after: %#v\n", sEmpty)

	arrB := [3]int{1, 2, 3}
	sB := arrB[1:3]
	// sB[0] = 22 // 在這邊會修改到arrB
	sB = append(sB, 4, 5)
	sB[0] = 22 // append後 在這邊不會修改到arrB
	fmt.Println(arrB)
	fmt.Println(sB)
}
