package main

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	Name    string
	Handle  RouteHandle
	Filters []Filter
}

type RouteHandle interface {
	AddRoute(url, method string, serverHandle ServerHandle, filters []FilterInterface)
	Match(context *Context) (Filter,error)
	http.Handler
}

type ServerHandle = func(c *Context) error

func (s *Server) Get(url string, handle ServerHandle, filters ...FilterInterface) {
	s.addRoute(url, "GET", handle, filters)
}

func (s *Server) Post(url string, handle ServerHandle, filters ...FilterInterface) {
	s.addRoute(url, "POST", handle, filters)
}

func (s *Server) Put(url string, handle ServerHandle, filters ...FilterInterface) {
	s.addRoute(url, "PUT", handle, filters)

}

func (s *Server) Delete(url string, handle ServerHandle, filters ...FilterInterface) {
	s.addRoute(url, "DELETE",handle, filters)
}

func (s *Server) addRoute(url string, method string, handle ServerHandle, filters []FilterInterface) {
	// 操作
	s.Handle.AddRoute(url, method, handle, filters)
}

func (s *Server) Run(addr string) {
	err := http.ListenAndServe(addr, s.Handle)
	if err != nil {
		panic(err)
		return
	}
}


func (s *Server) ShutDown(ctx context.Context)error {
	// 摘掉流量
	time.Sleep(50)
	return nil
}
// 自动检测是否实现了这个接口，没用其他意义
var _ RouteHandle = &BaseRouteHandle{}

func NewServer(name string, handle RouteHandle) *Server {
	return &Server{
		Name:   name,
		Handle: handle,
	}
}
