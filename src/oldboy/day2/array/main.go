package main

import (
	"fmt"
)

func main() {
	//声明 后 赋值
	var a [3]int
	a = [3]int{1, 2, 3}

	//直接赋值
	var b = [3]int{1, 2, 3}

	//...使用
	var c = [...]int{1, 2, 3, 4, 5}

	//根据索引值初始化
	var e [100]int
	e = [100]int{99: 1}

	fmt.Println(a, b, c, e)

	//遍历数组 for
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	//range
	for index, value := range a {
		fmt.Println(a[index], value)
	}

	//数组求和
	_sum := 0
	sum := [...]int{1, 3, 5, 7}

	/*for i:=0; i<len(sum);i++  {
		_sum += sum[i]
	}*/

	for _, value := range sum {
		_sum += value
	}
	fmt.Println(_sum)

	//二维数组 声明 后赋值
	var a2 [3][2]int
	a2 = [3][2]int{
		[2]int{
			1, 2,
		},
		[2]int{
			3, 4,
		},
	}

	fmt.Println(a2)
	//声明加赋值

	a3 := [...][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	fmt.Println(a3)

}
