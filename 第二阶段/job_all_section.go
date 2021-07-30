// 遍历映射中所有key组成的切片
package main

import "fmt"

func main () {
	// 1、遍历映射中所有key组成的切片
	user := map[string]int{"李浪":10,"刘强东":7,"马云":6,"任正非":9}

	keySlice := make([]string,len(user))			// 定义一个切片,长度为user的长度

	varSlice := []int{}			// 用map的值再定义一个切片

	i := 0
	for k, v := range user {			// 遍历索引,索引从keySlice中取
		keySlice[i] = k
		i++

		varSlice = append(varSlice,v)		// int切片可以使用append直接扩展元素，string切片不能，要定义len长度
	}
	fmt.Println(keySlice,varSlice)
}
