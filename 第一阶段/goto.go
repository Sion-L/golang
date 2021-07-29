package main

import "fmt"

// goto -- 跳转（保留c语言中的特性）
func main () {

	// 示例1：买包子和西瓜
	var yes string
	fmt.Print("有卖西瓜的吗？(Y/N): ")
	fmt.Scan(&yes)

	fmt.Println("小明的想法：")
	fmt.Println("十个包子")
	if yes != "Y" && yes != "y" {
		goto END			// 当输入的不是Y或者y的时候则跳转到END:的位置
	}
	fmt.Println("买一个西瓜")

	END:        // goto跳转位置的标签，用大写字母和:表示




	// 示例2：从1加到100
	sum := 0
	i := 1

	START:
	if i > 100 {
		goto FOEND    // i大于100的时候则跳转到FOEND结束循环，不然则执行下面的sum += i 和 i++，并跳转到START再次循环
	}
	sum += i
	i++
	goto START

	FOEND:
		fmt.Println(sum)



	// 示例3：两数相乘
	BREAKEND:
	for i := 0; i< 5; i++ {
		for j := 0; j < 5; j++ {
			if i*j == 9 {
				break BREAKEND // 如果i*j等于9，则跳出标签BREAKEND的for循环，即第一层for循环,跳出哪层循环标签写在哪层上面
			}
			fmt.Println(i,j,i*j)
		}
	}
}
