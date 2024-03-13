package main

import "fmt"

// ======= 实现层 ==========
var _ ComputerFactory = (*IntelFactory)(nil)

type IntelFactory struct{}

func (i *IntelFactory) CreateGPU() GPU {
	return new(IntelGPU)
}
func (i *IntelFactory) CreateRAM() RAM {
	return new(IntelRAM)
}
func (i *IntelFactory) CreateCPU() CPU {
	return new(IntelCPU)
}

type IntelGPU struct{}

func (i *IntelGPU) Display() {
	fmt.Println("我是intel显卡，有显示功能")
}

type IntelRAM struct{}

func (i *IntelRAM) Storage() {
	fmt.Println("我是intel内存,有存储功能")
}

type IntelCPU struct{}

func (i *IntelCPU) Calculate() {
	fmt.Println("我是intel的cpu，有计算功能")
}
