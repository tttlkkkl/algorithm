package main // import "github.com/tttlkkkl/algorithm"

import "fmt"

func main() {

	// link 链表逆序
	// link()
	var s = []int{2, 3, 4, 7, 3, 8, 9, 10, 99, 383, 0, 333, 1}
	fmt.Println("排序前：    ", s)
	// 插入排序
	InsertionSort(cp(s))
	// 选择排序
	SelectSort(cp(s))
	// 冒泡排序
	BubbleSort(cp(s))
}

func cp(s []int) []int {
	var a = make([]int, len(s), cap(s))
	copy(a, s)
	return a
}
