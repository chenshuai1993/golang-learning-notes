package main

import (
	"fmt"
	"os"
)

func main() {

	args()
}

// os.Args 字符切片
func args() {
	if len(os.Args) > 0 {
		for key, value := range os.Args {
			fmt.Printf("%d=>%v\n", key, value)
		}
	}
}
