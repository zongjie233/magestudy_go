package main

import "fmt"

type Plane struct {
	color string
}

func (Plane) fly() int {
	return 500
}

func (c Car) run() {
	fmt.Printf("run at %d km/h\n", c.speed)
}

type Bird struct {
	Name string
	Plane
	Car
}

func (Bird) fly() int {
	return 1000
}
