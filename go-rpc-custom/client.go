package go_rpc_custom

import (
	"fmt"
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
	// 通过反射获取fPtr未初始化的函数原型，获取fPtr指向的值的反射对象
	fn := reflect.ValueOf(fPtr).Elem()
	// 定义一个函数
	f := func(args []reflect.Value) []reflect.Value {
		// 处理请求参数
		reqArgs := make([]interface{}, 0, len(args))
		for _, item := range args {
			reqArgs = append(reqArgs, item.Interface())
		}
		// 编码
		reqData, err := GobEncode(Param{Name: name, Args: reqArgs})
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
		respData, err := GobDecode(respBytes)
		if err != nil {
			panic(err)
		}
		respArgs := make([]reflect.Value, 0, len(respData.Args))
		for i, arg := range respData.Args {
			// 必须进行nil转换，当值为nil，反射对象的类型是reflect.Value,但其有效性是 Invalid ，在使用某些时可能会产生panic
			if arg == nil {
				respArgs = append(respArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			respArgs = append(respArgs, reflect.ValueOf(arg))
		}
		return respArgs
	}
	// 输出函数类型信息
	fmt.Println("Function Name:", fn)
	fmt.Println("Function Type:", fn.Type())
	// 输出函数参数和返回值类型信息
	fmt.Println("Number of Input Parameters:", fn.Type().NumIn())
	for i := 0; i < fn.Type().NumIn(); i++ {
		fmt.Printf("Parameter %d Type: %v\n", i+1, fn.Type().In(i))
	}
	fmt.Println("Number of Output Parameters:", fn.Type().NumOut())
	for i := 0; i < fn.Type().NumOut(); i++ {
		fmt.Printf("Return Value %d Type: %v\n", i+1, fn.Type().Out(i))
	}
	// 参数1：一个函数类型，类型是reflect.Type
	// 参数2：一个函数，作用是对第一个函数参数操作，并返回结果，为reflect.value类型
	// MakeFunc 根据指定的函数类型创建一个新的函数，该函数的行为由参数2指定
	v := reflect.MakeFunc(fn.Type(), f)
	// 为函数fPtr赋值
	fn.Set(v)
}
