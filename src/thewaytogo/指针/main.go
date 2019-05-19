package main

import "fmt"

type Saiyan struct {
	name  string
	Power int
}

func main() {
	goku := &Saiyan{"Goku", 9000}
	Super(goku)
	fmt.Println(goku.Power)
}

/**
复制一个指针比复制一个复杂的结构的消耗小多了。在 64 位的机器上面，一个指针占据 64 bit 的空间。
如果我们有一个包含很多字段的结构，创建它的副本将会是一个很昂贵的操作。指针的真正价值在于能够分享它所指向的值
*/
func Super(s *Saiyan) {
	s.Power = 10
}
