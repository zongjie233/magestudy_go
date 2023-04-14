package main

type Fish interface {
	walk(length float32) error
	eat() error
}

type Crew interface {
	climb()
}

type qingwa struct {
	Name   string
	length float32
}

func (q qingwa) walk(a float32) error {
	return nil
}
func (q qingwa) climb() {

}
