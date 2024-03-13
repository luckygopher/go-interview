package main

import "fmt"

var _ ComputerFactory = (*KingstonFactory)(nil)

type KingstonFactory struct{}

func (i *KingstonFactory) CreateGPU() GPU {
	return new(KingstonGPU)
}
func (i *KingstonFactory) CreateRAM() RAM {
	return new(KingstonRAM)
}
func (i *KingstonFactory) CreateCPU() CPU {
	return new(KingstonCPU)
}

type KingstonGPU struct{}

func (i *KingstonGPU) Display() {
	fmt.Println("我是kingston显卡，有显示功能")
}

type KingstonRAM struct{}

func (i *KingstonRAM) Storage() {
	fmt.Println("我是kingston内存,有存储功能")
}

type KingstonCPU struct{}

func (i *KingstonCPU) Calculate() {
	fmt.Println("我是kingston的cpu，有计算功能")
}
