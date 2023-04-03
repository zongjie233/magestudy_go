package main

import "fmt"

func update_slice(arr []int) {
	arr[0] = 777
	arr[1] = 888
	arr = append(arr, 222, 333)
	fmt.Println("arr in update func--", arr)
}

func main() {
	/*
		s2 := make([]int, 5, 9)
		s2[0] = 1
		s2[1] = 1
		s2[2] = 1
		s2[3] = 1
		s2[4] = 1
		fmt.Printf("s2 before append: %p\n", s2)
		s1 := make([]int, 4, 6)
		s1[0] = 2
		s1[1] = 2
		s1[2] = 2
		s1[3] = 2
		s2 = append(s2, s1...)
		fmt.Printf("s2 after append: %p\n", s2)
		// s1的size与s2的cap相比，来判断是否发生内存变化
	*/
	arr := make([]int, 2, 4)
	update_slice(arr)
	fmt.Println("arr--", arr)
}
