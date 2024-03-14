package main

import "fmt"

/**
设计模式：代理模式
此案例来源某大佬,觉得恰到好处，拿来仅做理解笔记
*/

// BeautyWoman 抽象主题
type BeautyWoman interface {
	// MakeEyesWithMan 对男人抛媚眼
	MakeEyesWithMan()
	// HappyWithMan 和男人共度美好的时光
	HappyWithMan()
}

// Panjinlian 具体主题
type Panjinlian struct{}

func (p *Panjinlian) MakeEyesWithMan() {
	fmt.Println("潘金莲对本馆抛了个眉眼")
}

func (p *Panjinlian) HappyWithMan() {
	fmt.Println("潘金莲和本馆共度了浪漫的约会。。。")
}

// WangPo 中间的代理人，王婆
type WangPo struct {
	woman BeautyWoman
}

func NewProxyWangPo(woman BeautyWoman) BeautyWoman {
	return &WangPo{woman: woman}
}

func (w WangPo) MakeEyesWithMan() {
	w.woman.MakeEyesWithMan()
}

func (w WangPo) HappyWithMan() {
	w.woman.HappyWithMan()
}

// 业务逻辑
func main() {
	// 西门想找金莲，让王婆来安排
	wangpo := NewProxyWangPo(new(Panjinlian))
	// 王婆命令金莲抛媚眼
	wangpo.MakeEyesWithMan()
	// 王婆命令金莲和西门约会
	wangpo.HappyWithMan()
}
