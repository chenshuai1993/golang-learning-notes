package main

import (
	"fmt"
	"sort"
)

func main() {
	//练习题
	var arr = [...]int{3, 7, 8, 9, 1}

	//从数组中得到切片
	slice := arr[:]

	//用自带的sort包排序
	sort.Ints(slice)

	fmt.Println(slice)

}
