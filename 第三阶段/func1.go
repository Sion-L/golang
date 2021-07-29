package main

import (
	"fmt"
)

// 函数也可以定义给一个变量
func addc(a,b int) int {
	return a+b
}

// 格式化函数 - callback ，将传递的数据按照每行打印还是按照一行|分割打印
// 形参为可变参数的函数，函数给函数赋值
func print(callback func(...string), args ...string) { // 给可变函数传递可变参数args，func(...string)为函数类型的形参,可以接受任意多个字符串变量的函数
	fmt.Println("print函数输出:")
	callback(args...)
}

func list(args ...string) {		// 定义一个有参数为string的可变形参的函数list，跟函数print的一样
	for i, v := range args {
		fmt.Println(i,":",v)
	}
}

func main() {
	fmt.Println("%T\n", addc)
	var f func(int, int) int = addc 	// 把函数addc定义成变量f
	fmt.Println("%T\n", f)
	fmt.Println(f(1, 4))	 // 给变量f传参，则直接计算函数addc

	print(list,"A","C","E")		// 传递时会把list传递给callback，callback又调用list函数，解包后把ACE又传递给list函数的args
										// 不能直接println打印，因为list没有返回值


	// 匿名函数 -- 只需要使用一次的时候定义
	sayHello := func(name string) {
		fmt.Println("Hello",name)   // 跟把函数定义成变量差不多一样，这里只是直接使用函数的值类型进行初始化
	}
	sayHello("kk")
	sayHello("cc")


	// 定义匿名函数并且调用
	func (name string) {
		fmt.Println("Hi",name)
	} ("dashen")


	// 将匿名函数定义成一个变量并传递给另一个函数
	values := func(args ...string) {   // 接受任意string类型的变量
		for _, v := range args {		// 只取value，不取索引
			fmt.Println(v)
		}
	}
	print(values,"A","B","C")	  // 把values函数传递给print函数并打印


	// 在调用函数的时候调用一个匿名函数并传递进去
	print(func(args ...string) {
		for _, v := range args {
			fmt.Println(v,"\t")   // 按行打印
		}
		fmt.Println()
	},"A","C","E")
}

// 总结：将函数定义成一个变量，形参为可变参数的函数，函数给函数赋值，匿名函数（在main函数中定义）、定义匿名函数并调用、将匿名函数定义成一个变量并传递给另一个函数