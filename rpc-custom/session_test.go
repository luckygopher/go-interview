package rpc_custom

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

func TestSession_ReadWrite(t *testing.T) {
	// 定义监听IP和端口
	addr := "127.0.0.1:8080"
	// 定义传输的数据
	data := "hello word"
	notice := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 一个协程写入数据
	go func() {
		defer wg.Done()
		defer close(notice)
		// 创建tcp连接
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		// 通知 读取方 连接已开始监听（避免还未完全启动并监听端口，连接操作就进来了）
		notice <- struct{}{}
		// 如果没有连接到达，Accept()方法会一直阻塞直到有新的连接到达或者发生错误
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal(err)
		}
		// 写数据
		s := NewSession(conn)
		err = s.Write([]byte(data))
		if err != nil {
			t.Fatal(err)
		}
	}()
	// 一个协程读取数据
	go func() {
		defer wg.Done()
		<-notice
		// 与指定服务建立连接，并进行通信
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		// 读取数据
		s := NewSession(conn)
		result, err := s.Read()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(result))
	}()
	wg.Wait()
}
