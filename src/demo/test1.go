package main

import "fmt"

func main() {
	s1 := "a"
	s2 := "b"

	s3 := fmt.Sprintf("%s %s", s1, s2)

	fmt.Printf("%s", s3)
	m := make(map[int]bool, 100)
	fmt.Println(len(m))

	m[3] = true
	m[5] = false

	fmt.Println(m[3])
	fmt.Println(m[5])
	fmt.Println(m[4])
}
