package main

import "fmt"

func main()  {
	const NAME string = "kk"
	fmt.Println(NAME)
	// NAME = "woniu" 报错，const常量不允许再次赋值

	//省略类型
	const SS = "niubi"
	fmt.Println(SS)

	//定义多个常量（类型相同）
	const BB,CC = 12,13
	fmt.Println(BB,CC)

	//定义多个常量（类型不相同）
	const (
		B1 string = "wangwu"
		B2 int = 16
	)
	fmt.Println(B1,B2)

	//定义多个常量，省略类型
	const A1,A2 = 15,"lisi"
	fmt.Println(A1,A2)

	//定义多个常量时，如果后面的常量不定义类型和赋值，则取第一个常量的类型和值
	const (
		C3 int = 3
		C4
		C5
	)
	fmt.Println(C3,C4,C5)
}
//常量一般都用大写