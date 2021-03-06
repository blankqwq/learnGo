package main

import (
	"context"
	"fmt"
	"sync"
)

// hook

type Hook func(ctx context.Context) error

func BuildShutdownHook(services ...*Server) Hook {
	return func(ctx context.Context) error {
		wg := sync.WaitGroup{}
		wg.Add(len(services))
		c1 := make(chan struct{}, 1)
		// 执行
		for _, s := range services {
			go func(svr *Server) {
				// 等待结束
				err := svr.ShutDown(ctx)
				if err != nil {
					// 关闭错误
					fmt.Println()
				}
				// 退出成功,为什么要睡?
				wg.Done()
			}(s)
		}

		go func() {
			// 全部结束
			wg.Add(1)
			wg.Wait()
			// 通知
			c1 <- struct{}{}
		}()
		select {
			case <-c1:
				// 正常结束
				fmt.Printf("正常结束\n")
				return nil
			case <-ctx.Done():
				fmt.Printf("结束超过时间\n")
				return HookTimeOutErr{}
		}
	}
}
