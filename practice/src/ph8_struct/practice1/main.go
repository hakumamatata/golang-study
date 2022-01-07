package main

import "fmt"

type student struct {
	name  string
	grade float64
}

func printInfo(s student) {
	fmt.Println("Name:", s.name)
	fmt.Printf("Grade: %0.3f\n", s.grade)
}

func setInfo(s student) student {
	s.name = "Alonzo Cole"
	s.grade = 92.3

	return s
}

func setInfoByPointer(s *student) {
	s.name = "Alonzo Cole2"
	s.grade = 92.3222
}

func main() {
	var s student
	s = setInfo(s)
	printInfo(s)

	var s2 student
	setInfoByPointer(&s2)
	printInfo(s2)
}
