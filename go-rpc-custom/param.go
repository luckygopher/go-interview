package go_rpc_custom

import (
	"bytes"
	"encoding/gob"
)

// Param 数据结构体
type Param struct {
	Name string        // 方法名
	Args []interface{} // 参数
}

// GobEncode gob序列化
func GobEncode(data Param) ([]byte, error) {
	// 创建一个字节缓冲区
	buf := bytes.Buffer{}
	// 创建一个Gob编码器，并将数据序列化到字节缓冲区中
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// GobDecode gob反序列化
func GobDecode(data []byte) (Param, error) {
	result := Param{}
	buf := bytes.NewBuffer(data)
	// 创建一个Gob解码器，并从字节缓冲区中反序列化一个 Param 类型的值
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(&result); err != nil {
		return Param{}, err
	}
	return result, nil
}
