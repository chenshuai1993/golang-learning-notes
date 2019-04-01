package main

import "fmt"

func main() {

	//编译器自动推断
	//name := "world1"
	//name = "hello1"

	//a , b  := "97", "98"

	var info = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(info); i++ {
		fmt.Println(info[i])
	}

	//fmt.Println(&a,&b)
}
