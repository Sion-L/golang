package main

import "fmt"

func main() {
	// 打印乘法口诀表
	for line := 1; line <= 9; line++ {    // 定义行数为9，从1开始，每行+1
		for i := 1; i <= line; i++{			// 定义列数
			//fmt.Print(i," * ", line , " = " , i*line,"\t")   // 不换行，制表符打印
			 fmt.Printf("%d * %d = %-2d    ",i,line,i*line)   //也可printf输出,向左对齐移两位
		}
		fmt.Println()		// 输出一个空值 - 空格
	}
}
