package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(os.Getwd())
	if fin, err := os.Open("D:\\Goproject\\basicpkg\\iopkg\\text.txt"); err != nil {
		fmt.Println(err)
		return
	} else {
		defer fin.Close()
	}
}
