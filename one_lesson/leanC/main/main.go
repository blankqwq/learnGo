package main

import (
	"context"
	"fmt"
	"time"
)

var b bool

func t() {
	for {
		if b {
			fmt.Println("timeout1====")
			return
		}
		fmt.Println("running1")
	}
}

func t2(timeChan chan struct{}) {
	for {
		select {
		case <-timeChan:
			fmt.Println("timeout2====")
			return
		default:
			fmt.Println("running2")
		}
	}
}

func cpu(ctx context.Context) {
	for {
		go func(c context.Context) {
			select {
			case <-ctx.Done():
				fmt.Println("获取完成，memory获取退出")
				return
			default:
				time.Sleep(time.Second * 2)
				fmt.Println("memory信息获取成功")
			}
		}(ctx)
		select {
		case <-ctx.Done():
			fmt.Println("获取完成，cpu获取退出")
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("cpu信息获取成功")
		}
	}
}

func main() {
	go t()
	c := make(chan struct{})
	go t2(c)
	ctx, cancel := context.WithCancel(context.Background())
	go cpu(ctx)
	time.Sleep(time.Second * 3)
	b = true
	c <- struct{}{}
	cancel()
	time.Sleep(time.Second * 5)
}
