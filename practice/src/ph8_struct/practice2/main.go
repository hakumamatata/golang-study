package main

import "fmt"

type car struct {
	name     string
	topSpeed float64
}

func nitroBoost(c *car) {
	c.topSpeed += 50
}

type part struct {
	desc  string
	count int
}

func doublePack(p *part) {
	p.count *= 2
}

func main() {
	var mustang car
	mustang.name = "Mustang Cobra"
	mustang.topSpeed = 225
	nitroBoost(&mustang)
	fmt.Println(mustang.name)
	fmt.Println(mustang.topSpeed)

	var fuses part
	fuses.desc = "Fuses"
	fuses.count = 5
	doublePack(&fuses)
	fmt.Println(fuses.desc)
	fmt.Println(fuses.count)
}
