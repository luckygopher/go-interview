package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

// 定义一个名为 Shape 的接口，该接口包括两个方法：Area() 用于计算图形的面积，Perimeter()
// 用于计算图形的周长。实现两个结构体 Rectangle 和 Circle，它们均满足 Shape 接口。编写一个函数 printShapeInfo，
// 该函数接受一个实现了 Shape 接口的对象作为参数，并打印该形状的面积和周长。

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

// Area 面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter %f\n", s.Perimeter())
}

// 编写一个 Go 程序，实现一个多生产者单消费者（MPSC）模型，使用 goroutine 和 channel。
// 生产者从文件中读取形状信息（例如矩形的长和宽或圆的半径），并将形状对象发送到一个 channel 中。
// 消费者从该 channel 接收形状对象，并调用 printShapeInfo 函数打印形状的面积和周长。
//
// 预期效果：
// 面试者需要展示对并发编程的理解，确保多个生产者能够正确地和一个消费者协作，无数据丢失。

func producer(filename string, ShapeChan chan<- Shape, wg *sync.WaitGroup) {
	defer wg.Done()
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open file err:", err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		shapeType := parts[0]
		switch shapeType {
		case "rectangle":
			if len(parts) != 3 {
				fmt.Printf("data err")
				continue
			}
			width, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				continue
			}
			height, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				continue
			}
			ShapeChan <- Rectangle{width, height}

		case "circle":
			if len(parts) != 2 {
				fmt.Printf("data err")
				continue
			}
			radius, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				continue
			}
			ShapeChan <- Circle{radius: radius}
		default:
			fmt.Println("type not supported")
		}
	}
}

func consumer(shapeChan <-chan Shape, done chan<- struct{}) {
	for shape := range shapeChan {
		printShapeInfo(shape)
	}
	done <- struct{}{}
}

func main() {
	r := Rectangle{Width: 10, Height: 5}
	printShapeInfo(r)

	shapeChan := make(chan Shape)
	doneChan := make(chan struct{})
	var wg sync.WaitGroup
	files := []string{"rectangle.txt", "circle.txt"}
	for _, filename := range files {
		wg.Add(1)
		go producer(filename, shapeChan, &wg)
	}

	go consumer(shapeChan, doneChan)

	wg.Wait()
	close(shapeChan)

	<-doneChan
}
