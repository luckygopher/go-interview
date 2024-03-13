package main

/**
抽象工厂设计模式
*/

// ======抽象层=========

type ComputerFactory interface {
	CreateGPU() GPU
	CreateRAM() RAM
	CreateCPU() CPU
}

type GPU interface {
	Display()
}

type RAM interface {
	Storage()
}

type CPU interface {
	Calculate()
}

func main() {
	var intelCpt, nvidiaCpt, kingstonCpt ComputerFactory
	intelCpt = new(IntelFactory)
	nvidiaCpt = new(NvidiaFactory)
	kingstonCpt = new(KingstonFactory)
	// 一：
	intelCpt.CreateCPU().Calculate()
	intelCpt.CreateGPU().Display()
	intelCpt.CreateRAM().Storage()

	// 二：
	intelCpt.CreateCPU().Calculate()
	nvidiaCpt.CreateGPU().Display()
	kingstonCpt.CreateRAM().Storage()
}
