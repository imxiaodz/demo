package main

import "fmt"

// 将num的地址赋给指针ptr，并通过ptr去修改num的值
// 这是否就是slice[index]可以改变自身和底层数组值的原因呢？
func main() {
	var a int
	fmt.Println(a)  // 0
	fmt.Println(&a) // 0xc0000b2008
	var p *int
	p = &a
	*p = 20
	fmt.Println(a) // 20
}
