/**
创造切片的各种方式
**/
package main

import "fmt"

func main() {
	//1.声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}
	// 2.:=
	s2 := []int{}
	// 3.make()
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)
	// 4.初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	// 5.从数组切片
	arr := [7]int{1, 2, 3, 4, 5, 9, 11}
	var s6 []int
	// 前包后不包
	s6 = arr[2:4]
	fmt.Println(s6)
	fmt.Println(len(s6))
	// 切片的容量，从底层数组被切的位置开始到最后一个的元素个数之和
	fmt.Println(cap(s6))
}
