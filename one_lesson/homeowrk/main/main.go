package main

import (
	"errors"
	"fmt"
)

func fb(start int) int  {
	if start<2 {
		return 1;
	}
	return fb(start-2)+fb(start-1)
}

func add(arr []int,index,value int) ([]int,error) {
	len := len(arr)
	if index>len-1 {
		return nil,errors.New("out of range")
	}
	// 开始插入
	res := make([]int,len+1)
	current:=0
	for i:=range arr{
		if i==index {
			res[current] = value;
			current+=1
		}
		res[current] = arr[i]
		current+=1
	}
	return res,nil
}

func del(arr []int,index int) ([]int,error)  {
	len := len(arr)
	if index>len-1 {
		return nil,errors.New("out of range")
	}
	// 开始插入
	res := make([]int,len-1)
	current:=0
	for i:=range arr{
		if i==index {
			continue;
		}
		res[current] = arr[i]
		current+=1
	}
	return res,nil
}


func main()  {
	println(fb(4));
	// slice add
	origin := []int{1,2,3,4,5,6}
	fmt.Printf("%v\n", origin)
	origin,_ =add(origin,2,1)
	fmt.Printf("%v\n", origin)
	origin,_ =add(origin,3,1000)
	fmt.Printf("%v\n", origin)
	origin,_ =del(origin,2)
	fmt.Printf("%v\n", origin)

}
