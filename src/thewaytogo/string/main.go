package main

import "fmt"

func main() {
	stra := "the spice must flow"
	//stra = "哈"
	byts := []byte(stra)
	strb := string(byts)

	fmt.Println(byts)
	fmt.Println(strb)

	fmt.Println(len("椒")) //3字节
}
