package main

import (
	"fmt"
	"time"
)

func ticker() {
	tk := time.NewTicker(1 * time.Second)
	for i := 0; i < 3; i++ {
		<-tk.C
		fmt.Println(time.Now().Unix())
	}
	tk.Stop()
}

func main() {
	////layout := "2006-01-02 15：04：05" // 格式模板必须为这个格式
	//layout2 := "15-04-05"
	//nowStr := time.Now().Format(layout2)
	//
	//fmt.Println(nowStr)
	//if t, err := time.Parse(layout2, nowStr); err == nil { // 将一串字符串解析成相应格式
	//	fmt.Println(t.Hour(), t.Minute(), t.Second())
	//} else {
	//	fmt.Println(err)
	//}
	ticker()
}
