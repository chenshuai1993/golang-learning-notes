package main

import (
	"fmt"
	"time"
)

var counter = 0

func main() {

	for i := 0; i < 100; i++ {
		go incr()
	}
	time.Sleep(time.Millisecond * 10)
}

func process() {
	for i := 1; i < 100; i++ {
		fmt.Println(i)
	}
}

//use process
func useprocess() {
	fmt.Println("start")
	go process()
	time.Sleep(time.Millisecond * 10) // this is bad, don't do this!
	fmt.Println("done")
}

func incr() {
	counter++
	fmt.Println(counter)
}
