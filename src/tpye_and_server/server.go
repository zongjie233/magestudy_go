package main

import "net/http"

//接口是一组行为的抽象，尽量用接口，以实现面向接口编程。

type Server interface {
	Route(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttsServer struct {
	Name string
}

func (s sdkHttsServer) Route(pattern string, handlerFunc http.HandlerFunc) {
	//TODO implement me
	panic("implement me")
}

func (s sdkHttsServer) Start(address string) error {
	//TODO implement me
	panic("implement me")
}
