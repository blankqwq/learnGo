package main

import "fmt"

type Parent struct {

}

func (p *Parent)Name()string  {
	return "Parent"
}

func (p *Parent)Print()  {
	// 找到谁打印谁，先找同样的,找不到都不会找父级
	fmt.Println(p.Name())
}

type Son struct {
	*Parent
}

//func (s *Son)Print()  {
//	fmt.Println(s.Name())
//}

func (s *Son)Name()string  {
	return "Son"
}

func main()  {
	// typedef 别名
	type test = int
	s := &Son{}
	s.Print()
}
