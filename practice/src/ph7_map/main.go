package main

import "fmt"

func main() {
	var mapA map[string]int
	mapA = make(map[string]int)
	mapB := make(map[string]int)
	fmt.Println(mapA)
	fmt.Println(mapB)

	mapA["Mary"] = 77
	mapA["Andy"] = 88
	fmt.Println(mapA)

	mapC := map[string]string{
		"abc": "apple",
		"def": "banana",
	}
	mapD := map[string]string{}
	fmt.Println(mapC)
	fmt.Println(mapD)

	// 判別是否為零值
	valString, ok := mapD["hi"]
	fmt.Println(valString, ok) // if ok == false 就是還沒賦予值 可以拿此值做一些情況判斷

	// delete
	mapD["hi"] = "ok"
	valString, ok = mapD["hi"]
	fmt.Println(valString, ok)
	delete(mapD, "hi")
	valString, ok = mapD["hi"]
	fmt.Println(valString, ok)

	// for ... range
	mapE := map[string]string{
		"tw": "TAWAN",
		"jp": "JAPAN",
	}
	for key, value := range mapE {
		fmt.Println(key, value)
	}

}
