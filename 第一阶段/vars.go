package main

import (
	fmt2 "fmt"
)

// 函数外的变量可以不被使用
var version string = "1,0"

func main() {
	// 定义一个string类型的变量me
	/*
	变量名必须满足标识符命名规则
	1、必须由非空的unicode字符串组成，数字...
	2、不能以数字开头
	3、不能为go的关键字
	4、避免和go预定义标识符冲突，true/false
	5、驼峰
	6、标识符区分大小写
	 */
	var me string
	// 变量未赋值就打印，则自动赋予0值
	fmt2.Println(me)
	// 定义的变量必须被使用，否则报错
	me = "kk"
	fmt2.Println(me)
	fmt2.Println(version)

	// 定义同一类型的多个变量
	var name, user string = "Lilang", "dashen"
	fmt2.Println(name,user)

	// 定义不同类型的多个变量
	var (
		age int = 22
		height float64 = 1.72
	)
	fmt2.Println(age,height)

	// 也可直接赋值，会自动给你设置类型
	var (
		s = "aa"
		a = "42"
	)
	fmt2.Println(s,a)

	// 不用var，简短声明，必须在函数（此为main）内部声明,也可自动生成类型
	Theshy := true
	fmt2.Println(Theshy)

	var (
		e = "财"
		r = "富"
		t = "自"
		y = "由"
	)
	fmt2.Println(e,r,t,y)

	// 也可同时赋值多个
	 var (
	 	tt,cc,vv = "dd",33,true
	 )
	 fmt2.Println(tt,cc,vv)

	 // s和ss的值进行替换
	 ss := "bb"
	fmt2.Println(ss)
	s,ss = ss,s
	fmt2.Println(s,ss)

}
// 变量只能被声明一次