package main

/**
设计模式：抽象工厂
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
	// 组合一：
	intelCpt.CreateCPU().Calculate()
	intelCpt.CreateGPU().Display()
	intelCpt.CreateRAM().Storage()

	// 组合二：
	intelCpt.CreateCPU().Calculate()
	nvidiaCpt.CreateGPU().Display()
	kingstonCpt.CreateRAM().Storage()
}
