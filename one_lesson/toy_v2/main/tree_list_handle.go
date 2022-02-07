package main

import (
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
		if c, right := matchPathNode(current, v, nil); right {
			current = c
		} else {
			createSubTreeNode(current, pathArr[i:], method, handle)
			sort.Slice(current.children, func(i, j int) bool {
				return current.children[i].typeSort > current.children[j].typeSort;
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
	for _, v := range pathArr {
		// 匹配
		if c, right := matchPathNode(current, v, context); right {
			if val, ok := c.handle[method]; ok {
				return val, nil
			}
			return nil, NewMethodNotAllow(method)
		}
	}
	return nil, NewNotFound()
}

func matchPathNode(head *Node, path string, context *Context) (*Node, bool) {
	var extendRoute *Node
	for _, child := range head.children {
		// 避免用户乱输入
		if context != nil {
			if child.matchFunc(path, context) {
				return child, true
			}
		}else{
			//注册走静态
			if child.path==path && path!="*" {
				return child,true
			}
		}
	}
	return extendRoute, extendRoute != nil
}

func newTreeNode(path string, matchFunc RouteMatchHandle, typeSort int) *Node {
	return &Node{path: path, children: make([]*Node, 5), matchFunc: matchFunc, handle: make(map[string]ServerHandle, 4), typeSort: typeSort}
}

func newStaticTreeNode(url string) *Node {
	return newTreeNode(url, func(path string, c *Context) bool {
		return url == path && path != "*"
	}, StaticRouteSort)
}

func newRouteParameter(url string) *Node {
	// 闭包的妙用
	return newTreeNode(url, func(path string, c *Context) bool {
		// 路由参数匹配
		return url == path && path != "*"
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
	for _, v := range sub {
		if v == "*" {
			child = newRouteAll(v)
		} else {
			child = newStaticTreeNode(v)
		}
		current.children = append(current.children, child)
		current = child

		if v == "*" {
			break
		}
	}
	current.handle[method] = handle
}
