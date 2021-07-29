package main

import "fmt"

func main() {
	fmt.Println("continue:")
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue   // continue - 跳过本次循环,即跳过3
		}
		fmt.Println(i)
	}

	fmt.Println("break:")
	for i := 0; i < 5; i++ {
		if i == 3 {
			break   // break - 跳出循环，即到2就停止
		}
		fmt.Println(i)
	}
}
