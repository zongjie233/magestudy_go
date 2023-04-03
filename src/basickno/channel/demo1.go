package main

import "fmt"

func main() {

	ch := make(chan byte, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5

	v := <-ch
	fmt.Println(v)
	v = <-ch
	fmt.Println(v)

	close(ch)
	for ele := range ch {
		fmt.Println(ele)
	}
	fmt.Println("--------")

	fmt.Println(len(ch))
}
