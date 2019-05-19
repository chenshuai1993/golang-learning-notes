package main

import "fmt"

func main() {
	s := Saiyan{
		Person: &Person{
			name: "chenshuai",
		},
		power: 100,
	}
	s.Introduce()
	fmt.Println(s.Person.name)
	fmt.Println(s.name)
}

type Person struct {
	name string
}

func (p Person) Introduce() {
	fmt.Printf("i'm %s\n", p.name)
}

type Saiyan struct {
	*Person //隐式组合
	power   int
}
