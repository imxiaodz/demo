package main

import "fmt"

func test1() {
	var s1 = []int{1, 2, 3, 4, 5, 6}
	s2 := s1[0:3]
	fmt.Printf("s2: %v\n", s2)
	s3 := s1[3:]
	fmt.Printf("s2: %v\n", s3)
	s5 := s1[:]
	fmt.Printf("s2: %v\n", s5)
}

func test2() {
	var a1 = [...]int{1, 2, 3, 4, 5}
	a2 := a1[:3]
	fmt.Printf("a2: %v\n", a2)
}

func main() {
	test1()
	test2()
}
