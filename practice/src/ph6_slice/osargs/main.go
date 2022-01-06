package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:]) // 可以使用切片運算子 排除第一個檔名
}
