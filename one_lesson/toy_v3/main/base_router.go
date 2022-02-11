package main

import (
	"errors"
	"net/http"
)

type RouteItem struct {
	handle ServerHandle
	filter Filter
}

type BaseRouteHandle struct {
	Route Tree
}

func (h *BaseRouteHandle) Match(context *Context) (Filter, error) {
	return h.Route.Match(context.R.RequestURI,context.R.Method,context)
}

func (h *BaseRouteHandle)AddRoute(url,method string,handle ServerHandle,filters []FilterInterface) {
	filters = append(filters, &CloseRequestFilter{})
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


func (h *BaseRouteHandle)ServeHTTP(resp http.ResponseWriter, req *http.Request)  {
	context:=NewContext(resp,req)
	// 匹配路由
	handle,err := h.Match(context)
	if err!=nil {
		if errors.Is(err,NotFoundErrors{}) {
			context.HttpErr(http.StatusNotFound,err.Error())
		}
		if errors.Is(err,MethodNotAllow{}) {
			context.HttpErr(http.StatusMethodNotAllowed,err.Error())
		}
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