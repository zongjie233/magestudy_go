package main

import (
	"errors"
	"fmt"
)

func test1(num ...float64) (float64, error) {
	res := 1.0
	for _, i := range num {
		res *= i
	}
	if res != 0 {
		return 1.0 / res, nil
	} else {
		return 0, errors.New("分母不能为零")
	}
}

func test2(args ...float64) (float64, error) {
	firstnum := args[0]
	if firstnum == 0 {
		return 0, errors.New("分母不能为零")
	}
	if len(args) == 1 {
		return 1 / firstnum, nil
	}
	remain := args[1:]
	res, err := test1(remain...)
	if err != nil {
		return 0, err
	}
	return 1 / firstnum * res, nil

}
func main() {
	fmt.Println(test1(2, 5))
	fmt.Println(test1(2, 0))
	fmt.Println(test2(2, 1))
	fmt.Println(test2(2, 0))
}
