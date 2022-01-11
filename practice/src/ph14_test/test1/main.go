package main

import (
	"fmt"
	"ph14_test/test1/types"
)

func main() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", types.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", types.JoinWithCommas(phrases))
}
