package main

import "fmt"

func main() {
	/*s := Saiyan{
		Name:  "chenshuai",
		Power: 100,
	}*/

	//结构体方法
	//s.Super()
	//fmt.Printf("power is %d", s.Power)

	//构造函数
	//chen := NewSaiyan("chenshuai", 20)
	//fmt.Println(chen.Power)

	//new(T)
	newSaiyan()
}

//结构体
type Saiyan struct {
	Name  string
	Power int
}

//构造函数方法
func (s *Saiyan) Super() {
	s.Power += 10000
}

//构造函数
func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}

//new 尽管缺少构造器，Go 语言却有一个内置的 new 函数，使用它来分配类型所需要的内存。 new(X) 的结果与 &X{} 相同。
func newSaiyan() {
	goku := new(Saiyan)
	goku.Power = 1
	goku.Name = "22"
	// same as
	//goku := &Saiyan{}

	fmt.Println(goku)

	//不太推荐
	/*goku := new(Saiyan)
	goku.name = "goku"
	goku.power = 9001

	//vs
	//这一种值得推荐
	goku := &Saiyan {
		Name: "goku",
		Power: 9000,
	}*/
}
