package rpc_custom

import (
	"encoding/json"
	"fmt"
	"net"
	"reflect"
)

type Server struct {
	// 地址
	addr string
	// 服务端维护的函数名到函数反射值的map
	funcs map[string]reflect.Value
}

// NewServer 创建服务端对象
func NewServer(addr string) *Server {
	return &Server{
		addr:  addr,
		funcs: make(map[string]reflect.Value),
	}
}

// Register 服务端绑定注册方法
// @param fName 函数名，
// @param f 真正的函数
func (s *Server) Register(fName string, f interface{}) {
	if _, ok := s.funcs[fName]; ok {
		return
	}
	// map中没有值，则将函数的反射值添加到map
	fVal := reflect.ValueOf(f)
	s.funcs[fName] = fVal
}

// Run 服务端等待调用
func (s *Server) Run() {
	// 监听
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		fmt.Printf("监听 %s err:%v", s.addr, err)
		return
	}
	for {
		// 接收连接请求
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept err:%v", err)
			return
		}
		// 创建会话
		connSession := NewSession(conn)
		// 读取连接请求数据
		reqByteData, err := connSession.Read()
		if err != nil {
			fmt.Printf("read err:%v", err)
			return
		}
		// 对数据解码
		reqData := Param{}
		if err := json.Unmarshal(reqByteData, &reqData); err != nil {
			fmt.Printf("decode err:%v", err)
			return
		}
		// 根据函数名Name获取函数反射值
		f, ok := s.funcs[reqData.Name]
		if !ok {
			fmt.Printf("函数 %s 不存在", reqData.Name)
			return
		}
		// 解析遍历客户端的参数，放到一个数组中
		reqArgs := make([]reflect.Value, 0, len(reqData.Args))
		for _, arg := range reqData.Args {
			reqArgs = append(reqArgs, reflect.ValueOf(arg))
		}
		// 反射调用方法，传入参数
		result := f.Call(reqArgs)
		// 解析遍历执行结构，放到一个数组中
		respArgs := make([]interface{}, 0, len(result))
		for _, item := range result {
			respArgs = append(respArgs, item.Interface())
		}
		// 包装数据，返回给客户端
		respData := Param{reqData.Name, respArgs}
		// 编码
		respBytes, err := json.Marshal(respData)
		if err != nil {
			fmt.Printf("encode err:%v", err)
			return
		}
		// 返回连接响应数据
		err = connSession.Write(respBytes)
		if err != nil {
			fmt.Printf("session write err:%v", err)
			return
		}
	}
}
