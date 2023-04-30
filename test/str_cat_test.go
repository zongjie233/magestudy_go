package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// go test -v str_cat_test.go -run=Buffer 根据正则表达式运行单元测试
func TestBuffer(t *testing.T) {
	hello := "hello"
	golang := "golang"
	var buffer bytes.Buffer
	buffer.WriteString(hello)
	buffer.WriteString(",")
	buffer.WriteString(golang)
	fmt.Println(buffer.String())

}

const LOOP int = 100

// 基准测试运行：go test -bench=StrCat -run=^$ str_cat_test.go
func BenchmarkStrCatWithBuffer(b *testing.B) {
	hello := "hello"
	golang := "golang"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		buffer.Grow(LOOP * 12)
		for i := 0; i < LOOP; i++ {
			buffer.WriteString(hello)
			buffer.WriteString(",")
			buffer.WriteString(golang)
		}

	}

}

func BenchmarkStrCatWithOperator(b *testing.B) {
	hello := "hello"
	golang := "golang"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for i := 0; i < LOOP; i++ {
			str += hello + "," + golang
		}
	}
}

func BenchmarkStrCatWithJoin(b *testing.B) {
	hello := "hello"
	golang := "golang"
	arr := make([]string, LOOP*2)
	for i := 0; i < LOOP; i++ {
		arr = append(arr, hello)
		arr = append(arr, golang)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strings.Join(arr, ",")
	}
}
