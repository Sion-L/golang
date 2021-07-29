// 在go语言中参数传递默认均为值传递（形参为实参变量的副本），对于引用类型函数因其底层共享数据结构，所以在函数内可对引用类型数据修改从而影响函数外的变量信息
package main

import "fmt"

func changeInt(a int) {		// 函数形参为int类型
	a = 100
}

func changeSlice(s []int) {		// 函数形参为切片类型
	s[0] = 100
}

func changePoint(p *int) {
	*p = 100
}

func main() {
	num := 1
	changeInt(num)
	fmt.Println(num)		// 值类型，无影响

	nums := []int{1,2,3}	// 引用类型，值影响
	changeSlice(nums)
	fmt.Println(nums)

	changePoint(&num)	// 无影响，值类型
	fmt.Println(num)

}