package main

import "fmt"

//go 指针
func main() {
	a := 1
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) //a:1 ptr:0xc000098000
	fmt.Println(b)                     //0xc000098000
	fmt.Println(*b)                    //1
	fmt.Printf("%T\n", b)              //*int
}
