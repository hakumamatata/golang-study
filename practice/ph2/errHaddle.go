package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a str")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	var strBuilder strings.Builder

	// remove \n ...etc.
	input = strings.TrimSpace(input)

	// if large data can use WriteString
	strBuilder.WriteString(input + "HERE??" + "YOU KNOW")
	//strBuilder.WriteString("HERE?")

	fmt.Print(strBuilder.String())
}
