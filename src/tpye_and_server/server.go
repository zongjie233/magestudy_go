package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//接口是一组行为的抽象，尽量用接口，以实现面向接口编程。

type Server interface {
	Route(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
}
type signUpReq struct {
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
func (s *sdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
	//TODO implement me
	http.HandleFunc(pattern, handlerFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	//TODO implement me
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{}
}

func Signup(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}

	ctx := &Context{
		W: w,
		R: r,
	}
	err := ctx.ReadJson(req)
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}
	resp := &commonResponse{
		Data: 123,
	}
	respJson, err := json.Marshal(resp)
	if err != nil {
	}

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, string(respJson))
}

/*
一个结构体有某个接口的所有方法，他就实现了这个接口
方法接收器，遇事不决用指针
*/
