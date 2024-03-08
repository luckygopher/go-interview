package rpc_custom

import (
	"encoding/json"
	"net"
	"reflect"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

// Call 实现通用的RPC客户端
// 函数具体实现在server端，client只有函数原型
// @param name 函数名
// @param fPtr 指向函数原型
// xxx.Call("query", &query)
func (c *Client) Call(name string, fPtr interface{}) {
	// 通过反射获取fPtr未初始化的函数原型
	fn := reflect.ValueOf(fPtr).Elem()
	// 定义一个函数
	f := func(args []reflect.Value) []reflect.Value {
		// 处理请求参数
		reqArgs := make([]interface{}, 0, len(args))
		for _, item := range args {
			reqArgs = append(reqArgs, item.Interface())
		}
		// 编码
		reqData, err := json.Marshal(Param{
			Name: name,
			Args: reqArgs,
		})
		if err != nil {
			panic(err)
		}
		// 写入请求数据
		connSession := NewSession(c.conn)
		if err := connSession.Write(reqData); err != nil {
			panic(err)
		}
		// 读取响应数据
		respBytes, err := connSession.Read()
		if err != nil {
			panic(err)
		}
		// 解码
		respData := Param{}
		if err := json.Unmarshal(respBytes, &respData); err != nil {
			panic(err)
		}
		respArgs := make([]reflect.Value, 0, len(respData.Args))
		for i, arg := range respData.Args {
			// todo 必须进行nil转换 此处存在问题
			if arg == nil {
				respArgs = append(respArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			respArgs = append(respArgs, reflect.ValueOf(arg))
		}
		return respArgs
	}
	// 参数1：一个未初始化函数的方法值，类型是reflect.Type
	// 参数2：另一个函数，作用是对第一个函数参数操作，返回reflect.value类型
	// MakeFunc 使用传入的函数原型，创建一个绑定 参数2 的新函数
	v := reflect.MakeFunc(fn.Type(), f)
	// 为函数fPtr赋值
	fn.Set(v)
}
