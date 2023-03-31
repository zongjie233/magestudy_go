package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//接口是一组行为的抽象，尽量用接口，以实现面向接口编程。

type Server interface {
	Route(pattern string, handlerFunc func(ctx *Context))
	Start(address string) error
}
type signUpReq struct {

	// Tag
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}
type sdkHttpServer struct {
	Name string
}
type commonResponse struct {
	Data int
}

// Route 路由实现
func (s *sdkHttpServer) Route(pattern string, handlerFunc func(ctx *Context)) {
	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		ctx := NewContext(writer, request)
		handlerFunc(ctx)
	})
}

func (s *sdkHttpServer) Start(address string) error {
	//TODO implement me
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{}
}

func Signup(ctx *Context) {
	req := &signUpReq{}

	err := ctx.ReadJson(req)

	if err != nil {
		ctx.BadRequestJson(err)
	}
	resp := &commonResponse{
		Data: 123,
	}
	err = ctx.WriteJson(http.StatusOK, resp)

	if err != nil {
		fmt.Printf("写入响应失败: %v", err)
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
	}

	fmt.Fprintf(ctx.W, string(respJson))
}

/*
一个结构体有某个接口的所有方法，他就实现了这个接口
方法接收器，遇事不决用指针
*/
/*
这段代码定义了一个Server接口和一个sdkHttpServer结构体，sdkHttpServer结构体实现了Server接口中定义的两个方法：Route和Start。此外，还定义了一个commonResponse结构体和一个signUpReq结构体。

Server接口定义了两个方法：Route和Start。Route方法用于注册处理程序，它将一个URL模式和一个处理函数关联起来。Start方法用于启动HTTP服务器，它将服务器绑定到一个地址上并开始监听连接请求。

sdkHttpServer结构体是Server接口的实现。它有一个Name字段，但在代码中并没有用到。Route方法使用http.HandleFunc函数来将请求路由到相应的处理函数。Start方法使用http.ListenAndServe函数来启动HTTP服务器。

commonResponse结构体只有一个Data字段，用于存储响应数据。

signUpReq结构体包含三个字段：Email，Password和ConfirmedPassword。它们都被标记为JSON标签，以便在JSON编码和解码过程中正确地映射字段名和JSON属性名。

最后，还定义了一个Signup函数，它是一个处理函数，用于处理/user/signup路由。该函数使用ctx.ReadJson函数从请求中读取JSON数据，然后使用ctx.WriteJson函数将响应写回客户端。如果发生错误，它会使用fmt.Printf函数打印错误信息。

*/
