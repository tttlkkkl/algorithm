package main

import "fmt"

// SelectSort 选择排序算法
// 升序排序
// 遍历 每次选择最小的数放在排序位置
func SelectSort(s []int) {
	l := len(s)
	for i := 0; i <= l; i++ {
		// 寻找[i,l)区间里的最小值,并交换
		for j := i + 1; j < l; j++ {
			if s[j] < s[i] {
				// 交换
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	fmt.Println("选择排序后：", s)
}

// BubbleSort 冒泡排序
func BubbleSort(s []int) {
	l := len(s)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	fmt.Println("冒泡排序后：", s)
}

// OpBubbleSort 优化后的冒泡排序，已经确定有序的，不再进行比较
func OpBubbleSort(s []int) {
	l := len(s)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l-i; j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	fmt.Println("冒泡排序后：", s)
}

// InsertionSort 插入排序
func InsertionSort(a []int) {
	l := len(a)
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if a[j] > a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
	fmt.Println("插入排序后：", a)
}
