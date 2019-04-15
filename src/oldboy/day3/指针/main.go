package main

import "fmt"

/**
 &取地址
 *取值
地址: 内存地址	0x12321
指针: 指针是有类型的， 存储的是 某一类 变量类型的 内存 *int, 指针是只读的，安全指针
指针可以比较 == ， 可以做逻辑运算
*/
func main()  {
	/*var a int = 10

	fmt.Println(a)
	b := &a

	fmt.Println(b)
	fmt.Printf("%T\n",b)
	fmt.Printf("%d\n",*b)*/


	//指针的申明
	/*arr := [3]int{1, 2, 3}
	modifyArray(&arr)*/

	//引用类型 mark
	//切片[] map[]

	//值类型的 指针类型 用 new
	/**
	new： 用来初始化值类型指针的
	make: 切片 slice, map，chan
	 */
	var a = new(int)
	*a =  10
	println(a);
	println(*a);


	/**
	panic(恐慌) 和 recover（恢复）
	编译: 词法 语法分析
	 */
}


/**
指针使用场景
 */

func modifyArray(a *[3]int) {
	fmt.Println( (*a)[0] );
}


