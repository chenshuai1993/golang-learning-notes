package main

import "fmt"

type animal interface {
	move()
}

type dog struct{}

func (d dog) move() {
	fmt.Println("狗会动")
}

func main() {
	//值类型引用
	zhileixing()
}

//值类型
func zhileixing() {
	var x animal
	var wangcai = dog{} //旺财是dog类型
	x = wangcai         //x可以接收dog类型
	var fugui = &dog{}  //富贵是*dog类型
	x = fugui           //x可以接收*dog类型

	v, ok := x.(*dog)
	fmt.Println(ok)
	if ok {
		fmt.Println("ok")
	} else {
		fmt.Println(v)
	}
}
