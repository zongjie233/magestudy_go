package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// Context 为什么选择为结构体
type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) ReadJson(req interface{}) error {
	r := c.R
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// 要返回掉，不然就会继续执行后面的代码
		return err
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		// 要返回掉，不然就会继续执行后面的代码
		return err
	}
	return nil
}

func (c *Context) WriteJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = c.W.Write(respJson)
	return err
}

func (c *Context) OkJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}
func (c *Context) SystemErrorJson(resp interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, resp)
}
func (c *Context) BadRequestJson(resp interface{}) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}

// Context封装

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		R: request,
		W: writer,
	}

}

/*
这段代码看起来是在实现一个HTTP框架中的Context结构体和相关方法，主要实现了HTTP请求的读取、响应的写入以及错误码的设置。

其中，Context结构体封装了HTTP请求的ResponseWriter和Request，提供了ReadJson、WriteJson、OkJson、SystemErrorJson和BadRequestJson
等方法。这些方法都是通过ResponseWriter来实现HTTP响应，而ReadJson方法则是实现了HTTP请求中body的读取和反序列化JSON对象的功能。同时，NewContext
函数则提供了Context结构体的初始化方法。

这个Context结构体是一个典型的中间件设计模式的体现，它将HTTP请求的上下文信息封装到一个结构体中，将对HTTP请求和响应的处理封装到方法中，方便在框架中统一调用。
*/
