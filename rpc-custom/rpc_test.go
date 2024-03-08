package rpc_custom

import (
	"encoding/gob"
	"fmt"
	"net"
	"testing"
)

type User struct {
	Name string
	Age  float64
}

func queryUser(uid float64) (User, error) {
	user := make(map[float64]User)
	user[0] = User{Name: "ls", Age: 20}
	user[1] = User{Name: "zs", Age: 21}
	if val, ok := user[uid]; ok {
		return val, nil
	}
	return User{}, fmt.Errorf("用户不存在%f", uid)
}

func TestRPC(t *testing.T) {
	// 需要对interface{}可能产生的类型进行注册
	gob.Register(User{})
	addr := "127.0.0.1:8080"
	// 创建服务器
	svr := NewServer(addr)
	svr.Register("queryUser", queryUser)
	go svr.Run()
	// 客户端获取连接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	// 创建客户端
	cli := NewClient(conn)
	// 声明函数原型
	var query func(uid float64) (User, error)
	cli.Call("queryUser", &query)
	// 得到查询结果
	u, err := query(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(u)
}
