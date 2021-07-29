/*panic和recover函数：go语言提供panic和recover函数用于处理运行时错误，当调用panic抛出错误，中断原有的控制流程，
常用于不可修复性错误。recover函数用于终止错误处理流程，仅在defer语句的函数中有效，用于截取错误处理流程，recover
只能捕获到最后一个错误*/
package main

import "fmt"

// 使用recover捕获panic抛出的错误
func test() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	panic("error")
	return
}

func main() {
	// recover - 捕获panic的值，让程序继续执行
	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()

	// panic
	fmt.Println("main start")
	panic("error")			// 运行时抛出一个错误

	fmt.Println("over")

	// 打印抓取的错误信息
	err := test()
	fmt.Println(err)
}
