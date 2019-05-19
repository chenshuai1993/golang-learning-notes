package main

import "fmt"

// 接口是一种类型（抽象类型）
// 如果一个结构体都拥有某个接口中的全部方法，那么我们认为该结构体实现了该接口。
type Animal interface {
	Say()
}

type newint int

func (n newint) say() {
	fmt.Println("aaaa")
}

func main() {
	var a newint = 10
	a.say()
}
