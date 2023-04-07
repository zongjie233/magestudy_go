package main

import "fmt"

func f1(i interface{}) {
	if v, ok := i.(int); ok {
		fmt.Printf("i is int %d\n", v)
	} else if v, ok := i.(float32); ok {
		fmt.Printf("i is float32 %f\n", v)
	} else {
		fmt.Println("other type")
	}
}

func f2(i interface{}) {
	switch i.(type) {
	case int:
		v := i.(int)
		fmt.Printf("i is int %d\n", v)
	case float32:
		v := i.(int)
		fmt.Printf("i is float32 %f\n", v)
	default:
		println("other type")
	}
}

func main() {
	var i int
	f1(i)
	var i2 float32
	f1(i2)
}
