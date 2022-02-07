package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R  *http.Request
	RouteParameter map[string]interface{}
}

func NewContext(resp http.ResponseWriter, req *http.Request) *Context {
	return &Context{W: resp, R: req}
}

func (c *Context)ReadFromJson(data interface{}) error {
	body,err:= ioutil.ReadAll(c.R.Body)
	if err!=nil {
		return err
	}
	err = json.Unmarshal(body,data)
	if err!=nil {
		return err
	}
	return nil
}

func (c *Context)WriteToJson(data interface{},msg string,code int)  {
	res,_ :=json.Marshal(Com{
		Data: data,
		Msg:  msg,
		Code: code,
	})
	c.W.WriteHeader(200)
	c.W.Write(res)
}

type Com struct {
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
	Code int `json:"code"`
}

func (c *Context)HttpOk(data interface{},msg string)  {
	c.WriteToJson(data,msg,0)
}

func (c *Context)HttpErr(error int,msg string)  {
	c.W.WriteHeader(error)
	c.W.Write([]byte(msg))
}