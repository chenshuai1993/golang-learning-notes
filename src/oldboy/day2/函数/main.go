package main

import "fmt"

func main()  {

}

//函数demo
func f1()  {
	fmt.Println("hello")
}


//函数demo
func f2(name string, age int)  {
	fmt.Println("hello")
}

//函数demo
func f3(name , addr string)  {
	fmt.Println("hello")
}

//可变参数
func f4(names ...string)  {
	fmt.Println("hello")
}

//函数的返回值 无返回值
func f5()  {
	fmt.Println("hello")
}

//函数的返回值 一个返回值
func f6() int{
	fmt.Println("hello")
	return 1
}

//函数的返回值 多个返回值
func f7()  (int, int){
	fmt.Println("hello")
	return  1 ,2
}

//函数的返回值 命名返回值
func f8(a, b int)  (sum int, sub int){
	sum = a + b
	sub = a - b

	return sum, sub
}

//匿名函数
var demo = func() {
	fmt.Println("hello")
}







