// error接口：go语言通过error接口实现错误处理的标准模式，通过使用函数返回值列表中的最后一个值返回错误信息，将错误的处理交由程序员主动进行处理
package main

import (
	"errors"
	"fmt"
)

// 返回值定义错误类型 error,创建错误类型的值 - errors.New()函数、fmt.Errorf()函数
// 无错误返回nil，
func name(a,b int) (int, error) {    // 定义返回值打印错误信息
	if b == 0 {
		return 0, errors.New("管望明是我儿")		// 返回值为错误信息
	}
	return a / b, nil		// 正确返回值
}

func main() {
	fmt.Println(name(4, 2))
	fmt.Println(name(4,0))     // 除数不能为0，此处没打印出错误信息

	// 打印错误值
	if v, err := name(1 , 0); err == nil {	// err为nil则正确
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	// 使用fmt.Error打印错误信息
	a := fmt.Errorf("Error: %s","name by zero")    // 错误信息赋值给一个变量
	fmt.Println("%T","%v\n", a,a)



}
