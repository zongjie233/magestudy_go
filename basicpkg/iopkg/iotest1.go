package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func readTextFile() {
	if fin, err := os.Open("demo.go"); err != nil {
		return
	} else {
		defer fin.Close()
		reader := bufio.NewReader(fin)
		for {
			if line, err := reader.ReadString('\n'); err == nil {
				line = strings.TrimRight(line, "\n")
				fmt.Println(line)
			} else {
				if err == io.EOF {
					if len(line) > 0 {
						fmt.Println(line)
					}
				} else {
					fmt.Println(err)
				}
				break

			}

		}

	}
}

//func writeTextfile() {
//	if fout, err := os.OpenFile("text.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666); err != nil { // 如何不写O_APPEND,则默认重写
//		fmt.Println(err)
//		return
//	} else {
//		defer fout.Close()
//		writer := bufio.NewWriter(fout)
//		writer.WriteString("第一行\n")
//		writer.WriteString("第二行\n")
//		writer.Flush() // 强制清空缓存，把缓存中的内容写入磁盘，如果不加这行代码则以上写入无法生效
//	}
//
//}

// func walkDir(path string) error {
// 	if subFiles, err := ioutil.ReadDir(path); err != nil {
// 		return err
// 	} else {
// 		for _, file := range subFiles {
// 			fmt.Println(file.Name())
// 			if file.IsDir() {
// 				if err := walkDir(filepath.Join(path, file.Name())); err != nil { // 递归
// 					return err
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }

//func logger() {
//	log.Printf("%d", 65)
//
//	fout, err := os.OpenFile("my.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
//	if err != nil {
//		return
//	}
//	defer fout.Close()
//	logWriter := log.New(fout, "CHINA", log.Ldate|log.Lmicroseconds) // 定义输出格式
//	logWriter.Printf("%s\n", "zdqissb")
//}

func main() {
	//writeTextfile()
	readTextFile()
	//logger()
}
