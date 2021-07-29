package main

import "fmt"

func main() {
	// 索引 => 记录已经加到的n
	// 记录结果

	// 计算1加到100的值
	/*result += 1
	result += 2
	...
	result += 100
	fmt.Println(result)
	*/
	/*
		i => 1....100
		result += i
	 */

	// 初始化子语句； 条件子语句； 后置子语句
	result := 0
	for i := 1; i <= 100; i++{   // 初始值，循环条件，每次循环的任务
		result += i
	}
	fmt.Println(result)

	result = 0   // 此处不能临时赋值变量为0
	i := 1      // 定义两个初始子语句，加条件后，result每次加上i++的值
	for i <= 100 {
		result += i
		i++
	}
	fmt.Println(result)


	// go中for死循环的写法
	/*
	i = 0
	for {
		fmt.Println(i)
		i++
	}
	*/

	// for range的功能：数组、切片、映射、管道、字符串  -- 遍历
	desc := "我爱中国"
	for i,ch := range desc {
		fmt.Printf("%d %T %q\n",i,ch,ch)   // 查看遍历字符串的类型，输出的0、3、6、9为索引
	}

}
// 总结：for有1个表达式 - 一个初始子语句,for有2个表达式 - 两个初始子语句，for有0个表达式 - 死循环，range - 多功能遍历