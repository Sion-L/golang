// 数组是具有相同类型的数据项组成的一组长度固定的序列，数据项叫做数组的元素，数组的长度必须是非负整数的常量，长度也是类型的一部分
// 数组声明需要指定组成元素的类型以及存储元素的数量（长度）。
// 数组字面量
	// 指定数组长度[length]type{v1,v2,...,vlength}
	// 使用初始化元素数量推导数组长度：[...]type{v1,v2,...,vlength}
	// 对指定位置元素进行初始化：[length]type{im:vm,...,sin:in}
package main

import "fmt"

func main() {
	var nums [10]int     // 中括号表示数组的长度,int为数据项的类型
	var t1 [5]bool
	var t2 [3]string
	fmt.Printf("%T",nums)
	fmt.Println(nums)		// 没有数据项则输出10个0
	fmt.Println(t1)		// 输出5个false
	fmt.Printf("%q",t2)

	// 字面量表示方法
	nums = [10]int{1 ,2 ,3}  // {}赋予数组字面量，只给三个赋值，后面七个为0
	fmt.Println(nums)
	nums = [10]int{0: 1,9: 10} // 通过索引给指定位数赋值,位数是从0开始的，随意20的长度为0-19
	fmt.Println(nums)
	// 另一种数组字面量表示方法
	var test = [...]int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1}	// []根据{}的数据项去推导数组长度
	//nums = [...]int{1, 2, 3, 1, 2, 3}  //nums前面定义了数组长度时，[...]为数组nums定义字面量时必须满足十个长度，不然报错
	fmt.Println(test)


	// 数组的操作（运算）
	// 1、关系运算，对比两个数组是否相等。对比时两个数组长度要相等
	nums01 := [3]int{1, 3, 4}
	nums02 := [3]int{1, 3, 5}
	fmt.Println(nums01 == nums02)

	// 2、获取数组的长度
	fmt.Println(len(nums02))

	// 3、索引：0，1，2，3...len(array) - 1
	fmt.Println(nums02[0],nums02[2])          // 取数组nums02索引0和索引2的元素

	// 4、为数组索引赋予元素（值）
	nums02[0] = 1000
	fmt.Println(nums02)
	// for循环取数组的值
	for i := 0; i < len(nums02); i++ {
		fmt.Println(i, ":",nums02[i])
	}
	// 用for range遍历数组的元素
	for _, value := range nums02 {			// _不能舍弃，不然value表示为索引，而不是数组元素。且在for和if中有内置的作用域，因此不用特意声明value
		// fmt.Println(index, ":",value)
		fmt.Println(value)    // 当我们只要value变量的值的时候，不使用index，则把index替换成_，表示为空白标识符
		// _ := 50，这行代码没有意义，空白标识符赋值无意义


	// 5、多维数组
		// 例1
		var marrays [3][2]int   // 此句代码表示为数组marrays[3]里面长度为2的int类型数组。此种为多维数组 - 3个长度为2的数组
		fmt.Println(marrays)
		fmt.Println(marrays[0])			// 输出的第一个元素为一个长度为2的数组
		fmt.Println(marrays[0][0])		// 输出为长度为2的数组的第一个元素
		marrays[0] = [2]int{1,3}		// 修改多维数组marray里第一个长度为2的数组元素为1和3
		fmt.Println(marrays)
		marrays[1][1] = 1000		// 修改多维数组marray里索引(即第二个长度为2的数组)为1的数组的第1个索引的元素为1000
		fmt.Println(marrays)
		// 多维数组字面量
		marrays = [3][2]int{{1,3},{2,4},{5,6}}		// 为3个长度为2的数组赋值
		fmt.Println(marrays)
		// for range 遍历多维数组
		for _, v := range marrays {		// 向遍历出长度为2的数组，再取长度为2的数组的元素
			for _, vv := range v {
				fmt.Print(vv, "\t")
			}
			fmt.Println()
		}

		// 例2
		var m3 [3][2][5]int		// 三个元素为2长度均为5的数组。表示为3个元素，每个元素有2个数组的，2个数组长度均为5
		fmt.Println(m3)
	}

}