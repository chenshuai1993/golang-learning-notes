package 接口demo

import "fmt"

//interface
type personer interface {
	say(string)
	run(string)
}

//定义一个结构体
type man struct {
	name string
	age  int
}

//man.say()
func (m man) say(name string) {
	fmt.Println(name, "man-say")
}

//man.run()
func (m man) run(name string) {
	fmt.Println(name, "man-run")
}

func main() {
	//接口是一个类型
	/*var p personer

	m := man{
		name: "陈帅",
		age:  26,
	}

	p = m
	p.say(m.name);
	p.run(m.name);*/

	//m.say(m.name)
	//m.run(m.name)

	//空接口对象可以接受任意变量
	//_interface()

	//作为函数参数
	//useshow()

	//空接口作为map的值
	//usemap()

	//接口类型断言
	//assertion()

	//断言多次
	//var a interface{} = 1
	//justifyType(a)
}

//空接口可以变量 可以 接受任意变量
func _interface() {
	// 定义一个空接口x
	var x interface{}
	s := "Hello 沙河"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}

// 接口作函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

// 接口作为函数参数的例子
func useshow() {
	//空接口应用 做函数的参数
	show(6)
	show(6.66)
	show(true)
	show("hello")
	show([3]int{1, 2, 4})
	show([]int{1, 2, 4})
	show(map[int]string{1: "hh"})

	i := 1
	show(&i)

	type cat struct {
		name string
	}
	ccat := cat{name: "毛毛"}
	show(ccat)

}

/**
空接口作为map的值
*/
func usemap() {
	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}

//接口类型断言
func assertion() {
	var x interface{}
	x = "Hello 沙河"
	v, ok := x.(string)

	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}

//断言多次
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
