package main

import "fmt"




/**
main 函数
结构体 demo
*/
func main() {
	//自定义类型
	//UserDefinedType()

	//类型别名
	//typeAlias()

	//结构体
	//myStu()


	//结构体内存
	//mystructMem()

	//模拟构造函数
	/*my := myConstruct()
	fmt.Printf("%v\n",my)
	fmt.Println(my.a)*/



	//方法 和 方法接受者
	p := person{
		"chenshuai",
		27,
	}

	fmt.Println(p.name)
	p.say(p.name)

}


//方法和接受者
type person struct {
	name string
	age int
}

func (p person)say(name string)  {
	fmt.Println("你好",name)
}

/**
用户自定义类型
*/
func UserDefinedType() {
	type myint int

	var age myint = 10
	fmt.Printf("%T\n", age)
}

/**
类型别名
*/
func typeAlias() {
	type aliasInt = int

	var age aliasInt = 10
	var name  byte

	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", name)
}


/**
结构体
 */
func myStu()  {
	type addr struct {
		province string
		city string
	}

	type Stu struct {
		name string
		age int
		addr addr
	}

	//值类型 结构体变量
	var chenshuai = Stu {
		name: "陈帅",
		age:26,
		addr:addr{
			province: "河北",
			city:"石家庄",
		},
	}

	//指针类型结构体变量
	var chenshuaiPtr = &Stu {
		name: "陈帅",
		age:26,
		addr:addr{
			province: "河北",
			city:"石家庄",
		},
	}


	//只有实例化才会分配内存
	var chenshuai1 = Stu{}


	//创建指针类型结构体
	var zhizhen = new(Stu)
	zhizhen.name = "哈哈"


	fmt.Printf("%T\n", chenshuai)
	fmt.Printf("%T\n", chenshuaiPtr)
	fmt.Printf("%v\n", chenshuaiPtr)
	fmt.Printf("%#v\n", chenshuaiPtr)

	fmt.Println("name:", chenshuai.name)
	fmt.Println("age:", chenshuai.age)
	fmt.Println("province:", chenshuai.addr.province)
	fmt.Println("city:", chenshuai.addr.city)

	fmt.Println(chenshuai1.name)

	//展示详细信息
	fmt.Printf("%v\n", zhizhen)		//&{哈哈 0 { }}
	fmt.Printf("%#v\n", zhizhen)   //&main.Stu{name:"哈哈", age:0, addr:main.addr{province:"", city:""}}
	fmt.Printf("%s\n", zhizhen.name) 	//哈哈


	//new 初始化
	fmt.Printf("%s",chenshuaiPtr.name)  //chenshuaiPtr.name == (*chenshuaiPtr).name

	//使用键值对初始化
	var chenshuai_kv = Stu {
		"chenshuai",
		27,
		addr{
			"河北",
			"石家庄",
		},
	}

	fmt.Printf("%#v",chenshuai_kv)

}


/**
结构体内存 连续的一快内存
 */
func mystructMem()  {
	type test struct {
		a int8
		b int8
		c int8
	}

	n := test{
		1,2,3,
	}
	//%P 指针
	fmt.Print("\n")
	fmt.Printf("%p\n", &n.a)
	fmt.Printf("%p\n", &n.b)
	fmt.Printf("%p\n", &n.c)
	/**
	0xc0000160a8
	0xc0000160a9
	0xc0000160aa
	没一个int8 占用一个字节 ： 一个字节 = 1 batys = 8bit
	 */
}

// dsss
type test1 struct {
	a int8
	b int8
	c int8
}
/**
结构体构造函数
 */
func myConstruct() (*test1){

	//返回结构体值
	/*return test1{
		1,2,3,
	}*/

	//返回结构体值指针
	return &test1{
		1,2,3,
	}
}






