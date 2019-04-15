package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	//tojson 和 结构体tag
	type Stu struct {
		Name string `json:"name"`
		Age int	`json:"age"`
	}

	type classes struct {
		Title string
		Stus []*Stu
	}

	c := &classes{
		Title: "3年一班",
		Stus: []*Stu{},
	}

	for i := 0; i < 10; i++ {
		stu := &Stu{
			Name:   fmt.Sprintf("stu%02d", i),
			Age: i,
		}
		c.Stus = append(c.Stus, stu)
	}

	//tojson JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(data);
	fmt.Printf("json:%#s\n", data)



	/*fmt.Println(c.Stus)
	for index,v := range c.Stus {
		fmt.Println(index)
		fmt.Println(v.Age)
	}*/


	//反序列化 JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"3年一班","Stus":[{"name":"stu00","age":0},{"name":"stu01","age":1},{"name":"stu02","age":2},{"name":"stu03","age":3},{"name":"stu04","age":4},{"name":"stu05","age":5},{"name":"stu06","age":6},{"name":"stu07","age":7},{"name":"stu08","age":8},{"name":"stu09","age":9}]}`

	c1 := &classes{}

	err1 := json.Unmarshal([]byte(str), c1)
	if err1 != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}
