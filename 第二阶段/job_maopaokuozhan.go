package main

import (
	"fmt"
	"sort"
)

func main() {
	// 排序 - sort包
	nums := []int{1,3,5,9,8,7,4}
	sort.Ints(nums)
	fmt.Println(nums)


	// 二分查找 - 在有序的数组中查找（前提就是数组是有序的）
	fmt.Println(sort.SearchInts(nums,5))			// 数字5的索引是3
	fmt.Println(sort.SearchInts(nums,6))			// 数字6没有，但会打印出要插入6的索引位置

		// 判断一个数在数组中是否存在
	num := 5
	idx := sort.SearchInts(nums, 6)		// 取出数字5在数组中的索引
	if nums[idx] == num {		// 然后判断此索引的值是否为5
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}

}
