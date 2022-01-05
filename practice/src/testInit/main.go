package main

import (
	"fmt"
	_ "inits/bird"
	_ "inits/cat"
	_ "inits/dog" // 如果有依賴關係的會先執行
)

func init() {
	fmt.Println("main init...!")
}

func main() {
	fmt.Println("main main...")
}

/*
因為bird有先依賴dog 故雖然import順序有差別 但是dog會先執行
print:
dog bark!
bird bark!BallBall
cat bark!
cat bark! two!!
main init...!
main main...
*/
