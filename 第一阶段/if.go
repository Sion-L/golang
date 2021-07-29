package main

import "fmt"

// 小明：买十个包子，有卖西瓜的再买一个西瓜
// 小王：有卖西瓜的，只买一个包子
func main() {
	fmt.Println("小明的想法:")
	fmt.Println("十个包子")
	var yes string
	fmt.Println("有卖西瓜的吗?(Y/N)")
	fmt.Scan(&yes)   // scan函数种的变量要用&指定内存地址
	// yes == "Y" "y"   yes等于大y或小y
	if yes == "Y" || yes == "y" {     // 如果变量yes输入的值为Y或y，则输出买一个西瓜
		fmt.Println("买一个西瓜")
	}

	// if else
	fmt.Println("小王的想法:")
	if yes == "Y" || yes == "y" {
		fmt.Println("一个包子")
	} else {
		fmt.Print("十个包子")
	}

	// 多分支 -- if , else if
	var score int
	fmt.Print("请输入你的成绩：")
	fmt.Scan(&score)

	if score >= 90 {
		fmt.Println("评级：A")
	} else if score >= 80 {
			fmt.Println("评级：B")
		} else if score >= 70 {
				fmt.Println("评级: C")
			} else if score < 60 {
					fmt.Println("评级：D")
				}
}
