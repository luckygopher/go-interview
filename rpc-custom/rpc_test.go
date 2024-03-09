package rpc_custom

import (
	"encoding/gob"
	"fmt"
	"net"
	"testing"
)

type User struct {
	Name string
	Age  int
}

type MyErr struct {
	Msg string
}

func (m MyErr) Error() string {
	return m.Msg
}

func queryUser(uid int) (User, error) {
	users := make(map[int]User)
	users[0] = User{Name: "ls", Age: 20}
	users[1] = User{Name: "zs", Age: 21}
	if val, ok := users[uid]; ok {
		return val, nil
	}
	return User{}, MyErr{Msg: "用户不存在"}
}

func TestRPC(t *testing.T) {
	// 需要对interface{}可能产生的类型进行注册
	gob.Register(User{})
	gob.Register(MyErr{})
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
	var query func(int) (User, error)
	cli.Call("queryUser", &query)
	// 得到查询结果
	u, err := query(3)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(u)
}
