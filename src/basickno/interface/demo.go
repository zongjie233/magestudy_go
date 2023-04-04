package main

import (
	"errors"
	"fmt"
)

/* 接口是一组行为规范的集合
 结构体实现接口内所有函数，则称为实现了该接口
接口值由两部分组成，一个指向该接口的具体类型的指针，和一个指向该接口具体值的指针

interface {} 空接口可以接受任意类型的数据
*/

type Transporter interface {
	move(int, int) (int, error)
	whilst() error
}

type Car struct {
	price int
	power float32
}

func (c Car) move(a, b int) (int, error) {
	fmt.Println("I am moving")
	return a + b, nil
}

func (c Car) whilst() error {
	fmt.Println("I am whilsting")
	return errors.New("123")
}

func (c Car) walk() error {
	fmt.Println("I am walking")
	return nil
}

func main() {
	var ifc Transporter
	var c Car
	ifc = c
	ifc.whilst()
	ifc.move(3, 5)
	c.price = 90
	c.power = 900.0
	c.walk()
	c.move(5, 6)
}
