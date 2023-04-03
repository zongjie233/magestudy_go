package main

import (
	"time"
)

type PathError struct {
	path   string
	op     string
	opTime string
	msg    string
}

func NewPathError(path, op, msg string) *PathError {
	return &PathError{
		path:   path,
		op:     op,
		opTime: time.Now().Format("2006-01-02"),
		msg:    msg,
	}
}

// 自定义error必须实现一个方法

func (e *PathError) Error() string {
	return e.opTime + ":" + e.op + " " + e.path + " " + e.msg
}

func deletePath(path string) error {
	if 1 > 2 {
		//return errors.New("path not exist")
		return NewPathError(path, "delete", "path not exist")
	} else {
		return nil
	}

}

func main() {

}
