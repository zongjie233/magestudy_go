package main

import "fmt"

func main() {
	s2 := make([]int, 4, 7)
	s2[0] = 1
	s2[1] = 1
	s2[2] = 1
	s2[3] = 1
	fmt.Printf("s2 before append: %p\n", s2)
	s1 := make([]int, 4, 6)
	s1[0] = 2
	s1[1] = 2
	s1[2] = 2
	s1[3] = 2
	s2 = append(s2, s1...)
	fmt.Printf("s2 after append: %p\n", s2)
}
