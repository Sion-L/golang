package main

import "fmt"

// 同时为多个连续的参数定义类型
func addr(a,b int) int {
	return a+b
}

// 函数的可变参数 ...type,是一个切片，每个函数只能定义一个可变参数
func addN(a,b int, args ...int) int {
	total := a+b		// 两个参数的时候就是a+b
	for _, v := range args {   // 两个以上a+b后再每次加一个元素
		total +=v
	}
	return total
}

func cacl(op string, a, b int, args ...int) int {
	switch op {					// 传递一个add给op，则case
	case "add":
		return addN(a,b,args...)   // 给addN函数传值。在调用含有可变参数的函数时，需要对可变参数进行解包，不能直接传个数组过去
	}
	return -1		// 定义函数必须有返回值，且函数中由多个返回值时只运行第一个
}

// 定义多个返回值，类型可不一样
func test(a,b int) (int,int,int) {
	return a+b,a-b,a/b
}

// 命名返回值,类型一样可省略
func calc(a,b int) (sum,diff,product,merchant int) {
	sum = a + b
	diff = a - b
	product = a * b
	merchant = a / b
	return
}

func main() {
	fmt.Println(addr(1, 5))
	fmt.Println(addN(1, 5))
	fmt.Println(addN(1, 4, 5))
	fmt.Println(addN(1, 4, 5, 8))

	fmt.Println(cacl("add", 1, 2))
	fmt.Println(cacl("add", 1, 2, 5))

	fmt.Println(test(4,2))
	/*a,b,c := test(4,2)             也可再定义几个变量来接受返回值
	fmt.Println(a,b,c)*/

	fmt.Println(calc(10,5))     // 命名返回值结果

	// 拓展 - 取出切片中对应的元素，如取出除切片中第一个(1)元素的其它元素
	nums := []int{1, 2, 5, 8}
	fmt.Println(nums[:1])    // 第0个元素
	fmt.Println(nums[2:])
	nums = append(nums[:1],nums[2:]...)
	fmt.Println(nums)

}
// 总结：同种类型多参数定义、可变参数定义、可变参数解包、定义多个返回值、取数组中对应的元素值、接收返回值、命名返回值
