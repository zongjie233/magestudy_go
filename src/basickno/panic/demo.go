package main

import "fmt"

func test1() {
	fmt.Println("enter test1")
	defer func() {
		/*
				相当于
				err :recover()
				if err!=nil{}

			recover()用于捕获panic
		*/
		if err := recover(); err != nil { //只会捕获离他最近的panic
			fmt.Printf("里面发生了panic:%s\n", err)
		}
	}()
	fmt.Println("注册defer完成")

	defer func() {
		n := 0
		_ = 3 / n                   //引起panic，后续语句不执行
		fmt.Println("BBBBBBBBB")    // 不会执行
		defer fmt.Println("CCCCCC") // 不会执行
	}()
	defer fmt.Println("AAAAAAAAAAA")
	defer func() {
		arr := []int{1, 2}
		index := 6
		fmt.Println(arr[index])
	}()
}

func main() {
	//panic("1234565") //后边的语句都不会执行
	//fmt.Println("aaa")
	//recover会阻止panic继续执行
	test1()
}
