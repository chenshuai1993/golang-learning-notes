package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	//slice 扩容
	//sliceappend()

	//DEMO1
	//demo1()

	//copy
	_copy()

}

//DEMO
func demo() {
	scores := make([]int, 10, 20)
	scores[7] = 9033
	scores = append(scores, scores[:]...)
	fmt.Println(scores)
	//scores := make([]int, 10)
}

//slice 扩容
func sliceappend() {
	scores := make([]int, 0, 5)
	c := cap(scores)
	fmt.Println(c)

	for i := 0; i < 100000; i++ {
		scores = append(scores, i)

		// 如果容量改变了
		// Go 必须增加数组长度来容纳新的数据
		if cap(scores) != c {
			c = cap(scores)
			fmt.Println(c)
		}
	}
}

//demo 1
func demo1() {
	scores := make([]int, 5, 10)
	fmt.Println(scores)
}

//COPY
func _copy() {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	fmt.Println(scores)
	sort.Ints(scores)
	fmt.Println(scores)

	worst := make([]int, 5)
	copy(worst[2:4], scores[:5])
	fmt.Println(worst)
}
