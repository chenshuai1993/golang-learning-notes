package main

import (
	"fmt"
	"os"
)

func main() {
	agrs()
}

func agrs() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[1])
}
