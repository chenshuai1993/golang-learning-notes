package main

import "fmt"

func main() {
	prime();
}

//200~800质数
func prime()  {
	for i := 200; i < 800; i++ {
		flag := 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				flag++
			}
		}

		if flag == 1 {
			fmt.Println(i)
		}
	}
}
