package main

import "fmt"

//自定义类型
type newInt int

//go语言中 类型别名
//byte uint8 |  int int32


//类型别名
type myInt = int

//结构体
type student struct {
	name string
	age  int8
	gender string
	hobby []string
}


type mydemo struct {
	a string
	b string
}

//结构体
/**
std  Std 大写代表外部宝可访问  小写私有
 */
func std()  {
	var haojie = student{
		name : "11",
		age : 19,
		gender:"男",
		hobby: []string{"抽烟", "喝酒", "烫头"},
	}

	var haojie1 = student{
		name : "11",
		age : 19,
		gender:"男",
		hobby: []string{"抽烟", "喝酒", "烫头"},
	}

	//实例化 -基本实例化
	fmt.Printf("%v\n",&haojie.name)
	fmt.Printf("%v\n",&haojie1.name)
	fmt.Println(haojie.hobby)

	//实例化- new实例化
	//如果没有实例化对应的属性， 那么其属性值就是类型的默认属性值
	var wanzhan = student{}
	fmt.Printf("%v\n",wanzhan.name)

	//实例化- new实例化
	var yawei = new(student)
	yawei.name = "hello"
	yawei.age = 18
	fmt.Println(yawei.age)

	//初始化 - 只填写值初始化
	var stu = student{
		"姓名",
		20,
		"男",
		[]string{"抽烟", "喝酒", "烫头"},
	}
	fmt.Println(stu.name)


	//键值对初始化
	var stu1 = &student{
		name:"哈哈",
	}
	fmt.Println(stu1.name)
	fmt.Println(stu1.age)

}

//新类型 和 类型别名
func newtype()  {
	var  a  newInt = 10
	var  b  myInt = 10
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
}


//结构体的内存布局
//内存是以字节为单位的
//1字节 = 8位 = 8bit
func memory()  {
	type test struct {
		a int16
		b int8
		c int8
		d int8
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
	/**
	n.a 0xc00009e000   int8 占用一个字节
	n.b 0xc00009e001   一个字节
	n.c 0xc00009e002
	n.d 0xc00009e003
	 */
}


/**
结构体是一个值类型
 */
/*func testValuetype()  {*/
/*	  var */
/*}*/

func main()  {

	//新类型 和 类型别名
	//newtype();

	//结构体
	//std();

    //结构体内存布局
	memory()


}

