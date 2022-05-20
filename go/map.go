package main

import "fmt"

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 100
	scoreMap["王五"] = 100
	scoreMap["赵六"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	// 也可在声明时就填充元素
	userInfo := map[string]string{
		"username": "pprof.cn",
		"password": "123456",
	}
	fmt.Println(userInfo)

	// 遍历
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	// 只遍历key
	for k := range scoreMap {
		fmt.Println(k)
	}

	delete(scoreMap, "王五")
	for k, v := range scoreMap {
		println(k, v)
	}
}
