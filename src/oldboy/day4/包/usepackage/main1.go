package main

import (
	"fmt"

	"oldboy/day4/包/mypackage"
)

func init() {
	fmt.Println("我是 main 包的 init() ")

	fmt.Println("我是mypackage", mypackage.NAME)

}

func main() {
	mypackage.Hello()
}
