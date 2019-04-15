package main

import (
	"fmt"
	"strings"
)

func main()  {
	palindrome("how do you do")
}

func palindrome(string string)  {
	temp := strings.Split(string," ")
	//fmt.Println(temp)
	a := [3]int{1,2,3}
	b := []int{1,2,3}
	fmt.Printf("%T", temp)
	fmt.Printf("%T", a)
	fmt.Printf("%T", b)
	/*for key,val := range temp {
		fmt.Println(key,val)
	}*/
}
