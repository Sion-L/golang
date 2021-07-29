package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str = "hello 你好"

	// golang中string底层是通过byte数组实现的，所以直接求len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
	fmt.Println("len(str):", len(str))

	//以下两种都可以得到str的字符串长度

	//1、golang中的unicode/utf8包提供了用utf-8获取长度的方法
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))

	//2、通过rune类型处理unicode字符
	fmt.Println("rune:", len([]rune(str)))


	// golang中string底层是通过byte数组实现的。中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8
	// byte 等同于int8，常用来处理ascii字符，rune 等同于int32,常用来处理unicode或utf-8字符

}