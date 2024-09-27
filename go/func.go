package main

import "fmt"

func main() {
	v, e := multi_ret("one")
	fmt.Println(v, e) //输出 1 true

	v, e = multi_ret("four")
	fmt.Println(v, e) //输出 0 false

	//通常的用法(注意分号后有e)
	if v, e = multi_ret("one"); e {
		// 正常返回
		fmt.Println("1")
	} else {
		// 出错返回
		fmt.Println("0")
	}

	sum(1, 2)
	sum(1, 2, 3)
	//传数组
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	nextNumFunc := nextNum()
	for i := 0; i < 10; i++ {
		fmt.Println(nextNumFunc())
	}

	fmt.Println(fact(5))
}

func multi_ret(key string) (int, bool) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}

	var err bool
	var val int

	val, err = m[key]

	return val, err
}

func sum(nums ...int) {
	fmt.Print(nums, " ") //输出如 [1, 2, 3] 之类的数组
	total := 0
	for _, num := range nums { //要的是值而不是下标
		total += num
	}
	fmt.Println(total)
}

// 函数闭包
func nextNum() func() int {
	i, j := 1, 1
	return func() int {
		var tmp = i + j
		i, j = j, tmp
		return tmp
	}
}

// 函数递归
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
