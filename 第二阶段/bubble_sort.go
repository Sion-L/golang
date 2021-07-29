// 冒泡排序：https://www.cnblogs.com/bigdata-stone/p/10464243.html
// 从左到右每次两个数做比较，每次比较把最大的放到右边，最小的放在左边，如2，5，3; 2跟5比2小则不变，5跟3比3小则5放到最后，为2，3，5
// N个数字要排序完成，总共进行N-1趟排序，每i趟的排序次数为(N-i)次
// 俩个值交换的两种方法，1：tmp := a,a = b,b = tmp  2: a,b = b,a
package main

import "fmt"

func main() {
	/*heights := []int{6,7,5,9,10}
	for j := 0; j < len(heights); j++ {   // 比较次数为数组长度减一，即n个数经过n-1此排序才排序成功
		fmt.Println("第",j,"轮")
		for i := 0; i < len(heights)-1; i++ { // 从左到右每次两个数作比较
			if heights[i] > heights[i+1] { // 如果前面的数大于后面作比较的数则交换位置
				fmt.Println("交换：", heights[i], heights[i+1]) // 打印谁和谁交换
				tmp := heights[i]                            // 值的交换
				heights[i] = heights[i+1]
				heights[i+1] = tmp
			}
			fmt.Println(i, "交换完毕", heights) // 打印比较次数，并提示交换完成，输出交换后的数组
		}
		fmt.Println("第",j,"轮结果",heights)
	}*/

	// 优化，由于每次比较都会把最大的数放到随后，所以每次可以少比较一次
	heights := []int{6,7,5,9,10}
	for j := 0; j < len(heights); j++ {   // 比较次数为数组长度减一，即n个数经过n-1此排序才排序成功
		fmt.Println("第",j,"轮")
		for i := 0; i < len(heights)-1-j; i++ { // i定义索引数
			if heights[i] > heights[i+1] { // 如果前面的数大于后面作比较的数则交换位置
				fmt.Println("交换：", heights[i], heights[i+1]) // 打印谁和谁交换
				tmp := heights[i]                            // 值的交换
				heights[i] = heights[i+1]
				heights[i+1] = tmp
			}
			fmt.Println(i, "交换完毕", heights) // 打印比较次数，并提示交换完成，输出交换后的数组
		}
		fmt.Println("第",j,"轮结果",heights)
		}


}


