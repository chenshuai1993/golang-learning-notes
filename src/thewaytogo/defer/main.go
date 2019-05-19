package main

/**
尽管 Go 有一个垃圾回收器，一些资源仍然需要我们显示地释放他们。
例如，我们需要在使用完文件之后 Close() 他们。这种代码总是很危险。
一方面来说，当我们在写一个函数的时候，很容易忘记关闭我们声明了 10 行的东西。
另一方面，一个函数可能有多个返回点。Go 给出的解决方案是使用 defer 关键字：
*/
import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 读取文件
}
