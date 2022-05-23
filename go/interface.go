package main

import "fmt"

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵")
}

func (c cat) move() {
	fmt.Println("动")
}

func main() {
	var x animal
	x = cat{name: "花花"}
	x.move()
	x.say()
}
