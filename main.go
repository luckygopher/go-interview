package main

import "fmt"

func main() {
	// 题目：将6，2，10，32，9，5，18，14，30，29从小到大进行排列。
	// mpSort()
}

// 冒泡排序
func mpSort() {
	// 定义数组
	arr := [10]int{6, 2, 10, 32, 9, 5, 18, 14, 30, 29}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

// 选择排序
func xzSort() {

}

// 快排
func kpSort() {

}

// 堆排
func dpSort() {

}
