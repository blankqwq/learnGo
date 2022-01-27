package main

import (
	"errors"
	"net/http"
)


type RouteHandle interface {
	AddRoute(url,method string,serverHandle ServerHandle)
	Match(url string,method string) (error,ServerHandle)
	http.Handler
}


type BaseRouteHandle struct {
	Route map[string]map[string]ServerHandle
}

func (h *BaseRouteHandle)AddRoute(url,method string,handle ServerHandle) {
	h.Route[url] = make(map[string]ServerHandle)
	h.Route[url][method] = handle
}

func (h *BaseRouteHandle)Match(url string,method string) (error,ServerHandle) {
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
	err,handle := h.Match(context.R.RequestURI,context.R.Method)
	if err!=nil {
		return;
	}
	err = handle(context)
	if err != nil {
		return;
	}
	return
}