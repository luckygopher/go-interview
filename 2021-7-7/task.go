package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
)

func ChanTask() {
	// 题目：100个任务，每次执行10个，当err大于3时，取消任务
	task := make(chan struct{},100)
	ch := make(chan struct{},10)
	ctx, cancel := context.WithCancel(context.TODO())

	var errNUm int32
	defer close(ch)
	defer close(task)
	for i := 0; i < 100; i++ {
		ch <- struct{}{}
		go func() {
			if i > 100 {
				atomic.AddInt32(&errNUm,1)
				return
			}
			fmt.Println(i)
			<-ch
			task <- struct{}{}
		}()
		// 取消
		if errNUm >= 3 {
			cancel()
		}
	}
	for i := 0; i < 100; i++ {
		select {
		case <-task:
		case <-ctx.Done():
			break
		}
	}
}

// sync.waitGroup 示例
func WaitGroupDemo() {
	var sw *sync.WaitGroup
	urls := []string{
		"https://www.baidu.com",
		"https://www.jd.com",
	}
	for _, url := range urls {
		sw.Add(1)
		go func() {
			fmt.Printf("task start")
			defer sw.Done()
			_, err := http.Get(url)
			if err != nil {
				fmt.Printf("err %v",err)
			}
			fmt.Printf("task end")
		}()
	}
	sw.Wait()
}