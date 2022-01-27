package main

import "unicode/utf8"

func main(){
	a := "你好"

	println(a)
	println(len(a))

	// utf-8

	println(utf8.RuneCountInString(a))
}
