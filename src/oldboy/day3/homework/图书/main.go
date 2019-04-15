package main

import (
	"fmt"
	"os"
)

/**
	fmt.Println("1. 添加书籍")
	fmt.Println("2. 修改书籍信息")
	fmt.Println("3. 展示所有书籍")
	fmt.Println("4. 退出")
	fmt.Println()
 */

// 使用函数实现一个简单的图书管理系统。
// 每本书有书名、作者、价格、上架信息，
// 用户可以在控制台添加书籍、修改书籍信息、打印所有的书籍列表。

// Book is a struct
type Book struct {
	title   string
	author  string
	price   float32
	publish bool
}


//book的构造函数
func newBook(title, author string, price float32, publish bool)  (Book){
	return Book{
		title:title,
		author:author,
		price:price,
		publish:publish,
	}
}

func showmenu()  {
	fmt.Println("欢迎来到德莱联盟")
	fmt.Println("1: 添加书籍")
	fmt.Println("2: 修改书籍信息")
	fmt.Println("3: 展示所有书籍")
	fmt.Println("4: 退出")
}

func book(_input int32)  {

	if _input == 1 {

		var title string
		var author string
		var price float32
		var publish bool


		fmt.Println("请输入title:")
		fmt.Scanf("%s", &title)

		fmt.Println("请输入author:")
		fmt.Scanf("%s", &author)

		fmt.Println("请输入price:")
		fmt.Scanf("%f", &price)

		fmt.Println("请输入publish:")
		fmt.Scanf("%v", &publish)

		fmt.Println("你添加的书是:",title,author,price,publish)
		newbook := newBook(title,author,price,publish)
		booklist = append(booklist, newbook)

		book(0)

	}

	if _input == 3 {

		fmt.Println(booklist);

	}

	if _input == 4 {

		os.Exit(0)

	}





}


func main()  {
	//
	for {
		showmenu();
	}


	book(0)

	/*book := newBook("title", "hah", 10, false)
	fmt.Println(book.title)*/



}
