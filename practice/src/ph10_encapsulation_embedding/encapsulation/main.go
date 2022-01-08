package main

import (
	"fmt"
	"log"
	"ph10_encapsulation_embedding/calendar"
)

func main() {
	dateA := calendar.Date{}
	var err error
	err = dateA.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}
	err = dateA.SetMonth(10)
	if err != nil {
		log.Fatal(err)
	}
	err = dateA.SetDay(28)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dateA)
	fmt.Println(dateA.Year())
	fmt.Println(dateA.Month())
	fmt.Println(dateA.Day())
}
