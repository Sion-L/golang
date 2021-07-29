// 取数组中第二大的元素
package main

import "fmt"

func main() {
	// 方法一：冒泡排序两次，就会把数组中第二大的元素放在倒数第二位，取出可得

	// 方法二：第一个最大值和第二个最大值都初始化为0
	nums := []int{1,3,5,20,20,30,8,9,12}
	nummax,numsecond := nums[0],nums[0]
	for i, v := range  nums {
		if v > nummax { // 如果取出的元素比最大值还大，则使第二大的等于最大的,而最大的等于刚取出的最大的v
			numsecond = nummax
			nummax = v
		} else if  v ==nummax {		// 如果有多个则跳过
			continue
		} else if v > numsecond {		// 如果没有最大值，则判断v和第二大的，并当v大于第二大的时候，使第二大的等于v
			numsecond = v
		}
		fmt.Println(i,":",numsecond,numsecond)
	}
	fmt.Println(nummax,numsecond)
}
