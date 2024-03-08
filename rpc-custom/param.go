package rpc_custom

// Param json数据格式
type Param struct {
	Name string        `json:"name"` // 方法名
	Args []interface{} `json:"args"` // 参数
}
