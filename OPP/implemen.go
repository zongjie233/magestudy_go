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

func (b Bird) Fly() int {
	return 1000
}

type Fish interface {
	Swing()
	Breath()
}

type Frog struct {
	Name string
}

func (f Frog) Swing() {

}
