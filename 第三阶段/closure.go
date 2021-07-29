// 闭包，类似匿名函数，可以为函数的声明周期
package main

import (
	"fmt"
)

func main() {

// 子作用域可以调用父作用域的变量
	name := "kk"
	func() {
		fmt.Println(name)
	}()

//返回值定义成一个函数
	/*add1 := func(n int) int {
		return  n + 1
	}

	add2 := func(n int) int {
		return  n + 2
	}

	add3 := func(n int) int {
		return  n + 3
	}*/

		// 以上的add1，add2，add3可以整合在一起了
	addbase := func(base int) func(n int) int {
		return func(n int) int {		// 此处为返回一个函数
			return base + n		// 此处相当于函数里面调用了另外一个函数里面的变量
		}
	}
	add1 := addbase(1)
	add2 := addbase(2)
	add3 := addbase(3)
	add4 := addbase(4)
	fmt.Printf("%T\n",add4)  //
	fmt.Println(add4(5))
	fmt.Println(add1(3))
	fmt.Println(add2(3))
	fmt.Println(add3(3))
	fmt.Println(addbase(2)(3))
}