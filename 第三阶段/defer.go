// 延迟执行 - defer函数，可用来做错误处理，无论函数是否发生错误都在函数执行最后执行（return之前），若使用defer声明多个函数，则按照声明的顺序，先声明后执行
package main

import "fmt"

func main() {
	defer func() {		// defer关键字为当函数退出的时候执行
		fmt.Println("defer1")
	}()

	defer func() {
		fmt.Println("defer2")    // 当有多个defer时，按照先进后出，谁先占用谁后执行 - 栈的应用
	}()

	fmt.Println("main over")   // main over先打印
}