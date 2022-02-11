package main

import (
	"fmt"
	"sort"
	"strings"
)

type Tree struct {
	Node *Node
}

type RouteMatchHandle func(path string, c *Context) bool

const (
	StaticRouteSort = 999
	ParameterRoute  = 100
	WildCardRoute   = 0
)

type Node struct {
	path      string
	method    string
	matchFunc RouteMatchHandle
	handle    map[string]ServerHandle
	children  []*Node
	typeSort  int
}

func (t *Tree) createNodeHandle(path string, method string, handle ServerHandle) {
	path = strings.Trim(path, "/")
	pathArr := strings.Split(path, "/")
	current := t.Node
	for i, v := range pathArr {
		// 匹配
		if c, right := matchPathNode(current, v); right {
			current = c
		} else {
			createSubTreeNode(current, pathArr[i:], method, handle)
			sort.Slice(current.children, func(i, j int) bool {
				return current.children[i].typeSort > current.children[j].typeSort
			})
			return
		}
	}

}

func (t *Tree) Match(path, method string, context *Context) (ServerHandle, error) {
	// 若为 / 则
	println(path, method)
	path = strings.Trim(path, "/")
	pathArr := strings.Split(path, "/")
	current := t.Node

	if c, right := matchChild(current, pathArr, context); right {
		if val, ok := c.handle[method]; ok {
			return val, nil
		}
		return nil, NewMethodNotAllow(method)
	}
	return nil, NewNotFound()
}

func matchChild(head *Node, pathArr []string, context *Context) (*Node, bool) {
	fmt.Printf("%v\n", head.children)
	for i, v := range pathArr {
		for _, child := range head.children {
			if child.matchFunc(v, context) {
				if len(pathArr) == 1 {
					return child, true
				}
				return matchChild(child, pathArr[i+1:], context)
			}
		}
	}
	return nil, false
}

func matchPathNode(head *Node, path string) (*Node, bool) {
	var extendRoute *Node
	for _, child := range head.children {
		// 避免用户乱输入
		if child.path == path && path != "*" {
			return child, true
		}
	}
	return extendRoute, extendRoute != nil
}

func newTreeNode(path string, matchFunc RouteMatchHandle, typeSort int) *Node {
	// 初始容量为4,所以有nil,持续内存异常
	return &Node{path: path, children: make([]*Node, 0, 5), matchFunc: matchFunc, handle: make(map[string]ServerHandle, 4), typeSort: typeSort}
}

func newStaticTreeNode(url string) *Node {
	return newTreeNode(url, func(path string, c *Context) bool {
		return url == path && path != "*"
	}, StaticRouteSort)
}

func newRouteParameter(url string) *Node {
	// 闭包的妙用
	parameterName := url[1:]
	return newTreeNode(url, func(path string, c *Context) bool {
		// 路由参数匹配
		println(parameterName, path, "get!")
		if c != nil {
			c.RouteParameter[parameterName] = path
		}
		return true
	}, ParameterRoute)
}
func newRouteAll(url string) *Node {
	return newTreeNode(url, func(path string, c *Context) bool {
		// 路由参数匹配
		return true
	}, WildCardRoute)
}

func createSubTreeNode(current *Node, sub []string, method string, handle ServerHandle) {
	var child *Node
	flag := false
	for _, v := range sub {
		if v == "*" {
			child = newRouteAll(v)
			flag = true
		} else if v[0] == ':' {
			println("get parameter", v)
			child = newRouteParameter(v)
		} else {
			child = newStaticTreeNode(v)
		}
		current.children = append(current.children, child)
		current = child
		if flag {
			break
		}
	}
	current.handle[method] = handle
}
