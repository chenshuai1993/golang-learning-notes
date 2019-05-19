package main

import "fmt"

//数组 高效且死板

func main() {

	//定义
	var arr [10]int = [10]int{1, 2, 3, 4, 5}

	//[...]
	var a = [...]int{1, 3, 4, 5}

	fmt.Println(arr)
	fmt.Println(a)

	//遍历
	// for i,v := range arr {}
}
