package main

import "fmt"

// 闭包函数

func add(base int) func(int) int {
	return func(i int) int {
		fmt.Printf("base addr %p \n", &base)
		base += i
		return base
	}
}
func sub() func() {
	i := 10
	fmt.Printf("%p\n", &i)
	b := func() {
		fmt.Printf("i addr %p\n", &i)
		i--
		fmt.Println(i)
	}
	return b
}

func main() {
	b := sub()
	b()
	b()
	fmt.Println()

	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))

	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))
}
