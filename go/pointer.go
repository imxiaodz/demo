/*
每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。
Go语言中使用&字符放在变量前面对变量进行“取地址”操作
*/
package main

import "fmt"

func main() {
	a := 10
	b := &a                            // 取变量a的地址，将指针保存到b中
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)                    // 0xc00000e018
	c := *b                            // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}
