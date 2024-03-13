package main

import "fmt"

var _ ComputerFactory = (*NvidiaFactory)(nil)

type NvidiaFactory struct{}

func (i *NvidiaFactory) CreateGPU() GPU {
	return new(NvidiaGPU)
}
func (i *NvidiaFactory) CreateRAM() RAM {
	return new(NvidiaRAM)
}
func (i *NvidiaFactory) CreateCPU() CPU {
	return new(NvidiaCPU)
}

type NvidiaGPU struct{}

func (i *NvidiaGPU) Display() {
	fmt.Println("我是nvidia显卡，有显示功能")
}

type NvidiaRAM struct{}

func (i *NvidiaRAM) Storage() {
	fmt.Println("我是nvidia内存,有存储功能")
}

type NvidiaCPU struct{}

func (i *NvidiaCPU) Calculate() {
	fmt.Println("我是nvidia的cpu，有计算功能")
}
