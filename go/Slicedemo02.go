package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	fmt.Println(slice) //2,3,4

	// 地址一样
	fmt.Println(&arr[2])
	fmt.Println(&slice[1])

	// 改变了索引0的值，数组里的值也变了
	// slice[1] = 100     //这里可以看做*prt = 100 ?

	var p *int
	p = &arr[2]
	*p = 100
	fmt.Println(slice) //100,3,4
	fmt.Println(arr)   //1,2,100,4, 5
}
