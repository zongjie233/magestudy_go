package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name string
	Age  int  `place:"abc"`
	Sex  byte `json:"gender" xml:"sex"`
	Car
	height float32
}

type Car struct {
	speed int
}

func NewUser() *User {
	return &User{
		Name: "",
		Age:  -1,
		Sex:  3,
	}
}

// 需要有一个全局变量来保证只有一个实例
var (
	user     *User
	userOnce sync.Once
)

func GetUserInstance() *User {
	userOnce.Do(func() {
		if user == nil {
			user = NewUser()
		}
	})
	return user
}

func main1() {
	// 保证单例
	u1 := GetUserInstance()
	u2 := GetUserInstance()

	u3 := NewUser()
	u4 := NewUser()

	fmt.Printf("%p %p\n", u1, u2)
	fmt.Printf("%p %p", u3, u4)
}
