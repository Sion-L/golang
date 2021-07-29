// ----------切片
// 切片是长度可变的数组（具有相同数据类型的数据项组成的一组长度可变的序列），切片由三部分组成：
	// 指针：指向切片第一个元素指向的数据元素的地址
	// 长度：切片元素的数量
	// 容量：切片开始到结束位置元素的数量
	// 切片与数组的区别：切片不需要定义长度
// 声明：
	// 切片声明需要指定组成元素的类型，但不需要指定存储元素的数量(长度)，在切片声明后，会被初始化为nil，表示暂时不存在的切片
	// 默认情况下切片cap值小于1024的时候是成倍数增长的。比如有1变为2在变为4……但是超过1024后就变为原切片的1.25倍

package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums []int
	fmt.Printf("%T\n", nums)
	fmt.Println(nums == nil) // 由此判断切片的底层为指针
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))
	// 1、切片赋值 - 字面量
	nums = []int{1, 2, 3}
	fmt.Printf("%#v\n", nums)
	nums = []int{1, 2, 3, 4}
	fmt.Printf("%#v\n", nums)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))

	// 2、数组的切片得到的值也是切片，也能赋值给nums切片
	var array [10]int = [10]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	nums = array[1:10]
	fmt.Printf("%#v\n", nums)

	// 3、打印切片的长度和容量（len和cap）
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))

	// 4、make函数
	nums = make([]int, 3)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums)) // make函数可以给切片底层的指针赋值（对切片指针进行初始化）
	nums = make([]int, 3, 5)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))

	// 5、元素操作 - 增、删、改、查
	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(nums[2])
	nums[2] = 10
	fmt.Println(nums[2])
	// 增
	nums = append(nums, 1)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))	// append函数添加元素，打印发现元素+1
	nums = append(nums, 1)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))
	nums = append(nums, 1)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))	// 容量变为10了，数组的内存空间是连续的，当容量不够用，go会扩展容量
		// 遍历切片的值
	for index, value := range nums {
			fmt.Println(index, value)
	}
	nums = make([]int, 3, 10)		// 通过切片去获取一个新的切片

	// 3、切片的操作
	n := nums[1:3:8]
	// 容量8 - start1 = 8 - 1 = 7
	fmt.Printf("%T %#v %d %d\n",n,n,len(n),cap(n))
	// 容量3 - start2 = 10 - 2 = 8
	n = nums[2:3]
	fmt.Printf("%T %#v %d %d\n",n,n,len(n),cap(n))

	// 4、切片容量的问题
		// 当多个切片使用同一个数组时，append添加元素可能会使得某一数组的元素跟另一数组的元素冲突
	nums = make([]int,3,5)		// 构建一个
	nums01 := nums[1:3]
	fmt.Println(nums,nums01)
	nums01[0] = 1
	fmt.Println(nums,nums01)
	nums01 = append(nums01,3)
	fmt.Println(nums,nums01)
	nums = append(nums,5)
	fmt.Println(nums,nums01)
	nums = array[:]
	fmt.Println(nums,array)
	nums[0] = 100
	fmt.Println(nums,array)


	// 删除
		// 1、copy实现删除机制
	nums02 := []int{1,2,3}
	nums03 := []int{10,20,30,40}
	copy(nums03, nums02)   // 把nums02拷贝到nums03.目的在前，源在后
	fmt.Println(nums03)		// 拷贝的nums02的前三个元素会nums03的前三个元素覆盖

	nums03 = []int{10,20,30,40}
	copy(nums02,nums03)
	fmt.Println(nums02)			// 当把元素多的拷贝到元素少的，则只覆盖元素少的对应的元素数量

		// 2、删除第一个元素和删除最后一个元素,1~len 0~(len-1)
	nums04 := []int{1, 2, 3, 4, 5}
	fmt.Println(nums04[1:])      // 表示从1开始
	fmt.Println(nums04[:len(nums04)-1])   // 0可以省略，到len-1结束

		// 3、删除中间的元素，如nums04把3删除
	copy(nums04[2:],nums04[3:])    // 先把3:len copy 到2:len -> 即4,5 copy 3,4,5得到4,5,5
	fmt.Println(nums04[:len(nums04)-1])  // 然后从0开始到len，即1,2,4,5,5，再删除最后一个元素5，得到1，2，4，5 , 3被删除

		// 4、队列：每次添加在队尾，删除元素在队头（尾进头出，先进先出）
	queue := []int{}
	queue = append(queue,1)
	queue = append(queue,2)
	queue = append(queue,3)
	queue = append(queue,5)
	// 此时队列顺序为1，2，3，5
	fmt.Println(queue[0])
	queue = queue[1:]    // 把第一个索引0移除,尾进头出
	fmt.Println(queue)
	// 此时队列为2，3，5，即5进来1出去
	fmt.Println(queue[0])
	queue = queue[1:]
	fmt.Println(queue)  // 此时为3，5

		// 5、堆栈：每次添加在队尾，删除元素也在队尾（尾进尾出，先进后出）
	stack := []int{}
	stack = append(stack,1)
	stack = append(stack,2)
	stack = append(stack,3)
	fmt.Println(stack[len(stack)-1])   // 先打印最后一个元素
	stack = stack[:len(stack)-1]		// 再把最后一个元素移除，实现堆栈
	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]


	// 多维切片
	points := [][]int{}
	points01 := make([][]int,0)   // make指定切片的容量，跟上面的一样，0
	fmt.Printf("%T\n",points01)
	points = append(points,[]int{1,2,3})   // 添加元素
	points = append(points,[]int{3,4,0})
	points = append(points,[]int{3,5,7,9})
	fmt.Println(points)
	fmt.Println(points[0])
	fmt.Println(points[0:2])


	// 数组之间赋值
	slice01 := []int{1,2,3}
	slice02 := slice01
	slice01[0] = 10
	fmt.Println(slice02)   // 切片不是值类型，切片之间赋值指针、长度、元素、容量都会传递

	array01 := [3]int{1,2,3}
	array02 := array01
	array02[0] = 10	//
	fmt.Println(array01,array02)	// 数组是值类型，数组之间赋值是值传递


	// 切片的排序 - sort模块
	paixu := []int{4,5,7,8,6}
	sort.Ints(paixu)
	fmt.Println(paixu)		// 输出发现对切片的值进行了排序

	paixu01 := []string{"test","wo","ai","56","minzu"}
	sort.Strings(paixu01)
	fmt.Println(paixu01)	// 对字符串类型切片进行了排序，数字最小

	paixu02 := []float64{1.1,-1.2,1.3,1.5}
	sort.Float64s(paixu02)
	fmt.Println(paixu02)

	// 拓展
		// []int{} => 空切片 -> 底层的数组没有存储元素
		// var nums []int => nil切片 -> 底层数组的指针指向nil
}