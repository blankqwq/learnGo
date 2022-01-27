package main

import "net/http"

type Server struct {
	Name string
	Handle RouteHandle
}

type ServerHandle=func(c *Context) error

func (s *Server)Get(url string,handle ServerHandle)  {
	s.addRoute(url,"GET",handle)
}

func (s *Server)Post(url string,handle ServerHandle)  {
	s.addRoute(url,"POST",handle)
}

func (s *Server)Put(url string,handle ServerHandle)  {
	s.addRoute(url,"PUT",handle)

}

func (s *Server)Delete(url string,handle ServerHandle)  {
	s.addRoute(url,"DELETE",handle)
}

func (s *Server)addRoute(url string,method string,handle ServerHandle)  {
	// 操作
	s.Handle.AddRoute(url,method,handle)
}

func (s *Server)Run(addr string)  {
	err := http.ListenAndServe(addr, s.Handle)
	if err!=nil {
		println("error",err)
		return
	}
}

func NewServer(name string,handle RouteHandle) *Server  {
	return &Server{
		Name:   name,
		Handle: handle,
	}
}