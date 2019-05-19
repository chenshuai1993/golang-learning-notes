package main

import "fmt"

//map
func main() {
	//demo()

	//
	Saiyan()
}

func demo() {
	lookup := make(map[string]int)

	lookup["goku"] = 9001
	power, exists := lookup["goku"]

	// prints 0, false
	// 0 is the default value for an integer
	fmt.Println(power, exists)

	// returns 1
	total := len(lookup)
	fmt.Println(total)

	// has no return, can be called on a non-existing key
	delete(lookup, "goku")

	//映射是动态变化的。然而我们可以通过传递第二个参数到 make 方法来设置一个初始大小：
	//lookup1 := make(map[string]int, 100)

}

func Saiyan() {
	type Saiyan struct {
		Name    string
		Friends map[string]*Saiyan
	}

	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}

	goku1 := &Saiyan{
		Name:    "Goku1",
		Friends: make(map[string]*Saiyan),
	}
	goku.Friends["krillin"] = goku1

	//fmt.Println(goku)
	//fmt.Println(*goku.Friends["krillin"])

	lookup2 := map[string]int{
		"goku":  9001,
		"gohan": 2044,
	}
	fmt.Println(lookup2)

	map1 := map[int]int{}
	for i := 0; i < 10; i++ {
		map1[i] = i
	}

	//循环 map  迭代映射是没有顺序的。每次迭代查找将会随机返回键值对
	//无序
	for key, value := range map1 {
		fmt.Println(key, value)
	}
	/**
	8 8
	1 1
	...
	6 6
	*/
}
