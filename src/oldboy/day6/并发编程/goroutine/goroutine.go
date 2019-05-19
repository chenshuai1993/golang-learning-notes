package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //结构体 优雅的等待

func hello() {
	fmt.Println("hello hello()")
	wg.Done()
}

func main() {
	wg.Add(1)
	go hello()

	fmt.Println("hello main")
	wg.Wait() //在这里等待 go
}
