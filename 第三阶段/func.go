package main

import "fmt"

func sayHelloWorld() {          // 定义一个函数
	fmt.Println("Hello World!!")
}

func sayHi(name string) {			// 函数里面定义一个string类型的参数 - 即形参
	fmt.Println("你好:",name)
}

func add(a int,b int) int {		// 定义两个int类型的参数a和b，返回值类型也为int。返回值类型一般要用括号括起来，一个可以省略
	return a+b			// int返回值要通过return返回
}

func main() {
	sayHelloWorld()		// 调用函数
	fmt.Printf("%T\n",sayHelloWorld)		// 为func类型
	sayHi("lilang")		// 给函数中的参数传递值 - 即实参
	rt := add(1,5)		// 参数返回值为a+b，即1+5
	fmt.Println(rt)
}

// 总结：函数定义、函数调用、形参、实参、返回值类型、return返回
// 无参无返回值、有参有返回值、有参无返回值