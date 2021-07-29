package main

import "fmt"

// 选择 case
func main() {
	var yes string
	fmt.Print("有卖西瓜的吗？")
	fmt.Scan(&yes)

	fmt.Println("小明的想法")
	fmt.Println("十个包子")

	switch yes {
	case "y":      // 也可把下面的合并在一起,逗号分割,case "y","Y":
		fmt.Println("买一个西瓜-1")
	case "Y":
		fmt.Println("买一个西瓜-2")
	default:
		fmt.Println("买十个包子")   // y和Y没有匹配，则输出default
	}


	// switch多分支
	var score int
	fmt.Println("请输入你的成绩: ")
	fmt.Scan(&score)

	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}
