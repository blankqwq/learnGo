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
	Route map[string]map[string]Filter
}

func (h *BaseRouteHandle)AddRoute(url,method string,handle ServerHandle,filters []FilterInterface) {
	h.Route[url] = make(map[string]Filter)

	h.Route[url][method] = h.makeRouteFilter(handle,filters)
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

func (h *BaseRouteHandle)Match(url string,method string) (error,Filter) {
	println(url,method)
	if route,ok:=h.Route[url];ok {
		if h,ok:=route[method];ok  {
			return nil,h
		}else{
			return errors.New("method error"),nil;
		}
	}
	return errors.New("not found"),nil;
}

func (h *BaseRouteHandle)ServeHTTP(resp http.ResponseWriter, req *http.Request)  {
	context:=NewContext(resp,req)
	// 匹配路由
	err,handle := h.Match(req.RequestURI,req.Method)
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