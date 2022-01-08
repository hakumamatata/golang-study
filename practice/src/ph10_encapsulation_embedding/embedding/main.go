package main

import (
	"fmt"
	"log"
	"ph10_encapsulation_embedding/calendar"
)

func main() {
	// Event 崁入了 Date型別 可以使用匯出的屬性和方法
	eventA := calendar.Event{}
	err := eventA.SetTitle("abced")
	if err != nil {
		log.Fatal(err)
	}
	err = eventA.SetYear(2023)
	err = eventA.SetMonth(11)
	err = eventA.SetDay(30)
	fmt.Println(eventA.Title())
	fmt.Println(eventA.Year())
	fmt.Println(eventA.Month())
	fmt.Println(eventA.Day())
}
