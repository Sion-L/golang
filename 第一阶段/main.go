// main包
package main      //main包
import (
	"fmt"
)
/*
草泥马，亚视拉雷
 */
func main() {    // main主函数
	// 你吃屎阿里
	fmt.Println("hello world!")
	// Add()    这种方式为调用同文件夹下的其他go项目的Add函数
	// api.Baidu()      这种方式为调用其他文件夹下go项目中的Baidu函数
	// 注意，调用其他项目的函数时，函数名首字母要为大写，goland项目编译换成directory
	// 项目中函数名首字母为小写，表示此函数只能被当前包内部文件调用，首字母是大写，意味着任何包都能调用


	// 输出方式：内置函数和fmt包
	// 内置函数 -- print、println
	print("aa")
	println("bb")
	fmt.Print("cc")
	fmt.Println("dd")
	fmt.Println("ee","ff","tt")

	// fmt包 拓展：格式化输出
	// "潘科龙%s,吃屎拉。"      // %s为占位符，%d数字占位符
	var aa int = 3
	if aa > 3 {
		fmt.Println("error")
	}else {
		fmt.Println("access")
	}

}


