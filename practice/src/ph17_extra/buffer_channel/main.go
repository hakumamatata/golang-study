package main

import (
	"fmt"
	"time"
)

func printStr(c chan string) {
	time.Sleep(time.Second)
	c <- "a"
	time.Sleep(time.Second)
	c <- "b"
	time.Sleep(time.Second)
	c <- "c"
	time.Sleep(time.Second)
	c <- "d"
}

func main() {
	fmt.Println(time.Now())
	// 設定有緩衝的通道 可以不會在傳遞第一個值就鎖定
	// e.g. 設定為1時 "a" 和 "b" 會一起被執行
	//c := make(chan string)
	//c := make(chan string, 1)
	c := make(chan string, 3)
	go printStr(c)
	//time.Sleep(3 * time.Second)
	time.Sleep(5 * time.Second)
	fmt.Println(<-c, time.Now())
	fmt.Println(<-c, time.Now())
	fmt.Println(<-c, time.Now())
	fmt.Println(<-c, time.Now())
}
