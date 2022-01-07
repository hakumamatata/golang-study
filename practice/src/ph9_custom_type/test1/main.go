package main

import "fmt"

type Ll float64
type Gg float64

func main() {
	var LlA Ll
	var GgA Gg
	LlA = Ll(11.2)
	GgA = Gg(27.69)
	fmt.Println(LlA, GgA)

	// 兩個基於相同底層型別可以互相轉換 可能可以乘上倍率（例如　公升加侖）
	toGg := Gg(Ll(20.5) * 0.264)
	toLl := Ll(Gg(20.5) * 3.785)
	fmt.Println(toGg, toLl)

}
