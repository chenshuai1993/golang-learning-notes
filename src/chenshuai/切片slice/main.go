/**
切片删除操作
*/
package main

import "fmt"

func main() {

	//切片非空处理
	s1 := []string{"one", "", "there"}
	s1 = noEmpty(s1)
	fmt.Println(s1)

	s2 := noEmpty1(s1)
	fmt.Println(s2)

	//复制切片
	s3 := []int{1, 2, 3, 4, 5, 6}
	s4 := make([]int, 5)
	copy(s4[:], s3[0:5])

	fmt.Println(s3)
	fmt.Println(s4)
}

/**
desc:过滤非空
tip: 输入的切片 和 输出的切片 公用的一个底层切片， 底层切片被部分修改， 这样在函数内部就避免了重新分配一个数组
*/
func noEmpty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

/**
另外一种 复用同一个底层数组的方式
*/
func noEmpty1(strings []string) []string {
	out := strings[:0]
	for _, v := range strings {
		if v != "" {
			out = append(out, v)
		}
	}

	return out
}
