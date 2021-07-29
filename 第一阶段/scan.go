package main

import "fmt"
// 类c的scan，就像shell的read
func main() {
	var name string
	fmt.Print("请输入你的名称")
	fmt.Scan(&name)				// fmt包的scan函数,打印
	fmt.Println("你输入的名字是： ", name)

	var age int
	fmt.Print("请输入你的年龄")   // 定义一个整数类型的变量age，然后打印提示信息，后面再跟age变量的值，最后完整输出
	fmt.Scan(&age)  			// 注：如果输入的不是数字，则输出为0
	fmt.Println("你的年龄是： ",age)

	var height float64
	fmt.Print("请输入你的身高")
	fmt.Scan(&height)				// 如果中间的age类型输错，则height也默认输出为0
	fmt.Println("你的身高是；",height)
}
