package main

import "fmt"

//99乘法表
func main() {

	/*for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}*/

	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
	fmt.Printf("长度:%d, cap:%d", len(a), cap(a))
}
