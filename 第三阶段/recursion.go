// 递归是指函数直接或间接调用自己，递归常用于解决分治问题，把大问题拆分为相同的小问题进行解决，需要关注终止条件
// n的阶乘 - 例：4的阶乘!4 = 4x3x2x1=24
package main

import "fmt"

/*递归
func addN(n int) int {
	if n == 1 {			// 定义当n=1时（即为最小），则返回1
		return 1
	}
	return n + addN(n-1)
}

func main() {
	fmt.Println(addN(5))    // 结果为5+4+3+2+1=15
}*/

//阶乘
func factorial (n int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		return -1		// 一般-1表示异常
	}
	return  n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(5))  // 5x4x3x2x1=120
}
