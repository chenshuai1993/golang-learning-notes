package main

import "fmt"

func main() {

	//切面的声明1  声明
	var list = []int{1, 2, 3}
	for _, value := range list {
		fmt.Println(value)
	}

	//切片的声明2  数组得到
	var a = [3]int{1, 2, 3}
	var a1 = a[:]
	fmt.Printf("%T", a1)

	//切片的长度
	//len
	fmt.Printf("%d", len(a1))

	//切片的容量
	fmt.Printf("%d", cap(a1))
	fmt.Println()

	//扩容策略 x2

	//追加切片
	a1 = append(a1, 2)

	//copy()   make()

	//切片删除 go没有提供删除函数
	e := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%p\n", &e)

	x1 := e[:1]
	fmt.Printf("%p\n", &x1)

	e = append(e[:1], e[2:]...)
	fmt.Println("b的容量")

	//使用make 函数构造切片
	// make([]T, size, cap)
	makeQ := make([]int, 2, 10)
	makeQ = append(makeQ, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Println(makeQ)      //[0 0]
	fmt.Println(len(makeQ)) //2
	fmt.Println(cap(makeQ)) //10

	//append()方法添加切片元素
	makeQ = append(makeQ, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)

	//append()方法添加切片元素
	var citySlice []string
	// 追加一个元素
	citySlice = append(citySlice, "北京")
	// 追加多个元素
	citySlice = append(citySlice, "上海", "广州", "深圳")
	// 追加切片
	a2 := []string{"成都", "重庆"}
	citySlice = append(citySlice, a2...)
	fmt.Println(citySlice) //[北京 上海 广州 深圳 成都 重庆]

	//复制切片
	// copy()复制切片
	a3 := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 5)
	copy(c, a3)    //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1 2 3 4 5]
	c[0] = 1000
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1000 2 3 4 5]

	//从切片中删除元素 Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素
	// 从切片中删除元素 a = append(a[:index], a[index+1:]...)
	a4 := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a4 = append(a4[:2], a4[3:]...)
	fmt.Println(a4) //[30 31 33 34 35 36 37]

	//课后题
	var a5 = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a5 = append(a5, fmt.Sprintf("%v", i))
	}
	fmt.Println(a5)
	fmt.Printf("长度:%d, cap:%d", len(a5), cap(a5))

}
