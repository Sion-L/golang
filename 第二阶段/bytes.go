// 字节切片
package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 字节切片定义示例
	var test []byte = []byte{'a','b','c'}
	fmt.Printf(	"%T,%#v\n",test,test)


	// 类型转换
		// 1、字节切片转字符串,字符串再转换为字节切片
	s := string(test)
	fmt.Printf("%T,%#v",s,s)		// 输出为字符串abc
	bs := []byte(s)
	fmt.Printf("%T,%v\n",bs,bs)   // 字符串转换为字节切片
		//
	fmt.Println(bytes.Compare([]byte("abc"),[]byte("def")))		// 字符串转化为字节切片并比较
	fmt.Println(bytes.Index([]byte("abcdef"),[]byte("def")))	// 字符串转化为字节切片打印abcdef从def的d的索引
	fmt.Println(bytes.Contains([]byte("abcdef"),[]byte("abc")))		// 字符串转化为字节切片并判断abcdef是否包含abc

}