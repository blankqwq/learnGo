package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 1000; i++ {
		go func(n int) {
			for {
				fmt.Println(n)
				time.Sleep(1000)
			}
		}(i)
	}

	time.Sleep(time.Second * 100)
}
