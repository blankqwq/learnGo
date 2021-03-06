package main

import (
	"errors"
	"fmt"
	"strings"
)

type Tree struct {
	Node *Node
}

type Node struct {
	path string
	method string
	matchFunc func(path string,c *Context)
	handle map[string]ServerHandle
	children []*Node
}

func (t *Tree)createNodeHandle(path string,method string,handle ServerHandle)  {
	path = strings.Trim(path,"/")
	pathArr := strings.Split(path,"/")
	current := t.Node
	for i,v:=range pathArr {
		// 匹配
		if c,right:=matchPathNode(current,v,false);right {
			current = c
		} else{
			createSubTreeNode(current,pathArr[i:],method,handle)
			return;
		}
	}

}

func (t *Tree)Match(path,method string) (ServerHandle,error)  {
	// 若为 / 则
	path = strings.Trim(path,"/")
	pathArr := strings.Split(path,"/")
	current := t.Node
	for _,v:=range pathArr {
		// 匹配
		if c,right:=matchPathNode(current,v,true);right {
			if val,ok :=c.handle[method] ; ok {
				return val,nil
			}
			return nil,fmt.Errorf("method error")
		}
	}
	return nil,errors.New("not found")
}


func matchPathNode(head *Node,path string,isRegister bool) (*Node,bool)  {
	var extendRoute *Node
	for _,child := range head.children{
		// 避免用户乱输入
		println(child.path)
		if child.path == path && child.path!="*" {
			fmt.Println("get:%v",child)
			return child,true;
		}
		if child.path=="*" && isRegister {
			extendRoute = child
		}
	}
	return extendRoute,extendRoute!=nil;
}

func newTreeNode(path string) *Node  {
	return &Node{path: path,children: make([]*Node,5),handle: make(map[string]ServerHandle,4)}
}


func createSubTreeNode(current *Node,sub []string,method string,handle ServerHandle)  {
	for _,v:=range sub{
		child := newTreeNode(v)
		current.children = append(current.children,child)
		current = child
		if v=="*" {
			break
		}
	}
	current.handle[method] = handle
}