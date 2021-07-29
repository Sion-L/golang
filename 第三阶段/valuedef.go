// 值类型、引用类型
package main

import "fmt"

func main() {
	// 将变量赋给一个新的变量，并修改新变量的值，如果对旧变量有印象，则为引用类型，无影响则为值类型
	array := [5]string{"A","B","C"}
	slice := []string{"A","B","C"}

	arrayA := array
	sliceA := slice

	arrayA[0] = "C"
	sliceA[0] = "C"

	fmt.Println(array,arrayA)
	fmt.Println(slice,sliceA)

	// 值类型：int、bool、float、array、指针
	// 引用类型：slice、map


}
