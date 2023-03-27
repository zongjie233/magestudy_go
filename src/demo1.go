package src

import "fmt"

func demo() {
	slice := []int{1, 2, 3, 4, 5}
	subslice1 := slice[:3]
	subslice2 := slice[1:3]
	subslice3 := slice[2:3]

	fmt.Println(subslice1)
	fmt.Println(subslice2)
	fmt.Println(subslice3)
}
