package mypackage

import "fmt"

const NAME = "chenshuai"

func init() {
	fmt.Println("我 自定义 的包的 init() ")
}

func Hello() {
	fmt.Println("我自定义的包的输出")
}

func main() {
	fmt.Println("1111")
}
