/**
字符串翻转
*/
package main

import (
	"fmt"
	"strings"
)

func main() {

	//int 翻转
	//a := []int{0, 1, 2}
	//fanzhuanInt(a)
	//fmt.Println(a)

	//string 翻转
	//b := []string{"a", "b", "c", "d", "e"}
	//fanzhuanString(b)
	//fmt.Println(b)

	//翻转中英文
	c := []string{"a", "中", "c", "国", "e"}
	fanzhuanZY(c)
	fmt.Println(c)

	d := strings.ToLower("HELLO")
	fmt.Println(d)
}

//int 翻转
func fanzhuanInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//string 翻转
func fanzhuanString(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//翻转中英文混合
func fanzhuanZY(s []string) {
	_len := len(s) / 2
	for i, _ := range s {
		if i < _len {
			j := len(s) - 1 - i
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
}
