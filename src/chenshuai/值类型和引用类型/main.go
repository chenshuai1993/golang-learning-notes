/**
这是一个demo
*/
package main

import "fmt"

func main() {
	var a int
	var b float32
	var c bool
	var d string

	var e [3]int         //数组
	var f []int          //slias 切片
	var g map[int]string //map hash+无序集合
	var h *int           //

	//值类型
	type persoon struct {
		name string
		age  int
	}

	var p persoon

	//标准类型 值类型 int float bool sting [...]int 数组  变量声明后， 默认初始化会有改类型的零值  int 0, float 0, bool false, string ''， 数组对应标准类型的零值
	//复合类型 引用类型， 变量声明， 没有默认初始化变量初始化， 需要make()才初始化，才分配内存
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e[0])

	fmt.Println(f) //变量声明 &&
	x := make([]int, 3, 10)
	fmt.Println("------", x[1])

	fmt.Println(g)
	fmt.Println(g[10])

	fmt.Println(h == nil)

	fmt.Println("p-name=>", p.name)
	fmt.Println("p-age=>", p.age)

}
