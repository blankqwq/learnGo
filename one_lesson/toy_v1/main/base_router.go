package main

import (
	"net/http"
)

type RouteItem struct {
	handle ServerHandle
	filter Filter
}

type BaseRouteHandle struct {
	Route Tree
}

func (h *BaseRouteHandle)AddRoute(url,method string,handle ServerHandle,filters []FilterInterface) {
	h.Route.createNodeHandle(url,method,h.makeRouteFilter(handle,filters))
}

func (h *BaseRouteHandle) makeRouteFilter(handle ServerHandle,filters []FilterInterface) Filter  {
	var root Filter = func(c *Context) error {
		return handle(c)
	}
	for i := len(filters)-1; i > -1 ; i-- {
		f:=filters[i]
		root = f.Handle(root)
	}
	return root
}

func (h *BaseRouteHandle)Match(url string,method string) (Filter,error) {
	println(url,method)
	return h.Route.Match(url,method)
}

func (h *BaseRouteHandle)ServeHTTP(resp http.ResponseWriter, req *http.Request)  {
	context:=NewContext(resp,req)
	// 匹配路由
	handle,err := h.Match(req.RequestURI,req.Method)
	if err!=nil {
		context.HttpErr(http.StatusNotFound,err.Error())
		return;
	}
	err = handle(context)
	if err != nil {
		context.HttpErr(http.StatusInternalServerError,err.Error())
		return;
	}
	return
}

func NewBaseHandle() *BaseRouteHandle  {
	return &BaseRouteHandle{Tree{Node: &Node{}}}
}