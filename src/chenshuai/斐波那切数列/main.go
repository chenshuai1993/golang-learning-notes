package main

import "fmt"

func main() {
	r := fibo(10)
	fmt.Println(r)
}

//菲波那切数列
func fibo(n int) int {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

	return x
}
