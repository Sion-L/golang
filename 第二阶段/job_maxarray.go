// 找出数组中最大的元素
package main

import "fmt"

func main() {
	// 方法一：用冒泡，排序一次，最大值排在最后，得值

	// 方法二：初始值为数组得第0个元素
	nums := []int{1,3,20,30,5,8,9,11}
	maxnumx := nums[0]		// 第一个最大值为数组第一个元素 - 初始值
	for i, v := range nums {	// 遍历数组的值与maxnumx做比较，如果v大，则把他两交换。i为索引，v为元素 - 值
		if v > maxnumx {
			maxnumx = v
		}
		fmt.Println(i,":",maxnumx)
	}
	fmt.Println(maxnumx)


}
