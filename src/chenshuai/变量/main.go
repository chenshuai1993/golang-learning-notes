package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/*var age int8 = 127
	//fmt.Println(^age)
	fmt.Printf("%b\n", age)
	fmt.Printf("%b\n", age+1)
	fmt.Printf("%d\n", unsafe.Sizeof(age))

	var a int16
	var b int16

	var c string

	var d bool


	e := [3]int{1,2,3}
	f := [5]int{1,2,3,4,4}



	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &d)
	fmt.Printf("%p\n", &c)

	fmt.Printf("%d\n", unsafe.Sizeof(a))
	fmt.Printf("%d\n", unsafe.Sizeof(b))
	fmt.Printf("%d\n", unsafe.Sizeof(c))
	fmt.Printf("%d\n", unsafe.Sizeof(d))
	fmt.Printf("%d\n", unsafe.Sizeof(e))
	fmt.Printf("%d\n", unsafe.Sizeof(f))*/

	//string 大小
	str()

}

//字节大小
func zijie() {
	var i int = 1
	var ii int = 2
	fmt.Printf("%p\n", &i)
	fmt.Printf("%p\n", &ii)
	fmt.Printf("%d\n", unsafe.Sizeof(i))
	fmt.Printf("%d\n", unsafe.Sizeof(ii))

}

func str() {
	name := "h"
	j := "j"
	fmt.Printf("%p\n", &name)
	fmt.Printf("%p\n", &j)
	fmt.Printf("%d\n", unsafe.Sizeof(name))
	fmt.Printf("%d\n", unsafe.Sizeof(j))
	fmt.Printf("%T\n", name)
}

func strcopy() {
	//[i:j]  左边等于 右边不等 【i,j)

}
