package go_rpc_custom

import (
	"encoding/binary"
	"io"
	"net"
)

// 自定义协议 header 4字节 + data []byte

// Session 会话连接的结构体
type Session struct {
	conn net.Conn
}

// NewSession 创建连接
func NewSession(conn net.Conn) *Session {
	return &Session{conn: conn}
}

// 向连接中写数据
func (s *Session) Write(data []byte) error {
	// header 4字节表示数据长度 + 数据长度的切片
	buf := make([]byte, 4+len(data))
	// 写入头部数据，记录数据长度
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	// 写入数据
	copy(buf[4:], data)
	_, err := s.conn.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

// 从连接中读取数据
func (s *Session) Read() ([]byte, error) {
	// 读取头部长度
	header := make([]byte, 4)
	// 按头部长度 读取头部数据
	_, err := io.ReadFull(s.conn, header)
	if err != nil {
		return nil, err
	}
	// 读取数据长度，利用了ReadFull不会读取流已经被读取的部分，是从当前位置开始读取指定长度数据
	dataLen := binary.BigEndian.Uint32(header)
	// 按照数据长度去读取数据
	data := make([]byte, dataLen)
	_, err = io.ReadFull(s.conn, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
