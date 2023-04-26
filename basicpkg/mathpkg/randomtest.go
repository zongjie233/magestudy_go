package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const seed int64 = 43
	souce := rand.NewSource(seed)
	rander := rand.New(souce)

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", rander.Intn(10))
	}
	fmt.Println("\n")
	souce.Seed(seed) // 重置种子
	rander2 := rand.New(souce)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", rander2.Intn(10))
	}
	arr := rand.Perm(10)
	fmt.Println(arr)

	fmt.Println("\n")
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println(arr)
}
