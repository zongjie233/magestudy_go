package main

import "fmt"

func basic() {
	fmt.Println("A")
	defer fmt.Println(1)
	fmt.Println("B")

	// 后注册的先执行
	defer fmt.Println(2)
	fmt.Println("C")
}

func goo() (c int) {
	c = 999
	defer func(s string) {
		fmt.Printf("c = %d, %s\n", c, s)
	}("hs")
	// 后注册先执行
	defer fmt.Printf("c = %d\n", c)

	return 10
}

func main() {
	basic()
	goo()
}
