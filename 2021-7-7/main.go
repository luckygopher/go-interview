package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
)

func main() {
	// 题目：将6，2，10，32，9，5，18，14，30，29从小到大进行排列。
	// mpSort()
	WaitGroupDemo()
}

// 冒泡排序
func mpSort() {
	// 定义数组
	arr := [10]int{6, 2, 10, 32, 9, 5, 18, 14, 30, 29}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

// 选择排序
func xzSort() {

}

// 快排
func kpSort() {

}

// 堆排
func dpSort() {

}

func ChanTask() {
	// 题目：100个任务，每次执行10个，当err大于3时，取消任务
	task := make(chan struct{}, 100)
	ch := make(chan struct{}, 10)
	ctx, cancel := context.WithCancel(context.TODO())

	var errNUm int32
	defer close(ch)
	defer close(task)
	for i := 0; i < 100; i++ {
		ch <- struct{}{}
		go func() {
			if i > 100 {
				atomic.AddInt32(&errNUm, 1)
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
	var sw sync.WaitGroup
	var errNum int32 = 0
	ctx,cancel := context.WithCancel(context.Background())
	urls := []string{
		"https://www.baidu.com",
		"https://www.jd.com",
		"https://www.baidu.com",
		"https://www.jd.com",
		"htp://www.zhenaiwanghhh.com",
		"h://www.zhenaiwanghhh.com",
		"ht://www.zhenaiwanghhh.com",
		"https://www.jd.com",
	}
	for _, url := range urls {
		sw.Add(1)
		go func(ctx context.Context, errNum *int32, url string) {
			fmt.Printf("task start \n")
			defer sw.Done()
			fmt.Println(*errNum)
			if *errNum >= 2 {
				cancel()
			}
			go func(ctx context.Context) {
				select {
				case <-ctx.Done():
					fmt.Printf("cancel task")
					return
				}
			}(ctx)
			_, err := http.Get(url)

			if err != nil {
				fmt.Printf("err %v \n", err)
				atomic.AddInt32(errNum,1)
				return
			}
			fmt.Printf("task end \n")
		}(ctx, &errNum, url)
	}
	sw.Wait()
}

func ChanDemo() {

}
