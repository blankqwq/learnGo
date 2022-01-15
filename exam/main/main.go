package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main()  {
	var n,m int
	_, _ = fmt.Scanf("%d", &n)
	_, _ = fmt.Scanf("%d", &m)
	nums := getInt(n)
	if n<3 {
		fmt.Println(0)
		os.Exit(0)
	}
	res:=0
	// 去除重复项
	data := make(map[string]int,100)
 	for i:=0;i<n;i++{
		for j:=i+1;j<n-1;j++ {
			count := nums[i]+nums[j]+nums[j+1]
			if count%m==0{
				// 如何去除重复
				temp := []int{nums[i],nums[j],nums[j+1]}
				sort.Sort(sort.IntSlice(temp))
				tempStr := makeIntStr(temp,"-")
				if _,ok:=data[tempStr];ok {
					continue
				}
				data[tempStr]=1
				res +=1
			}
		}
	}
	fmt.Print(res)
}

func makeIntStr(data []int,str string) string  {
	res:=make([]string,len(data))
	for i,_ := range res{
		res[i]=strconv.Itoa(data[i])
	}
	return strings.Join(res,str)
}

func getInt(len int) []int {
	nums:=make([]int,len)
	for i,_:=range nums{
		_, _ = fmt.Scanf("%d", &nums[i])
	}
	return nums
}