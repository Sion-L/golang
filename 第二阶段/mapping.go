/* 映射，类似java的hashmap
      定义：映射是存储一系列无序的key/value对，通过key来对value进行操作（增、删、改、查）
	key/value规则
		映射的key只能为可使用==运算符的值类型（字符串、数字、布尔、数组），value可以为任意类型
 */
package main

import (
	"fmt"
)

func main() {
	var scores map[string]int	// key和value都是有类型的，[]指明key的类型，后面的int指明value的类型。
	fmt.Printf("%T %#v\n",scores,scores)		// %#v可打印所有类型的值，输出为nil的映射，映射是nil不能进行任何操作
	fmt.Println(scores == nil)
	scores = make(map[string]int)	// 也可用make函数声明映射，此句代码含义跟map[string]int{}一样，为空的映射
	fmt.Println(scores)

	// 1、字面量
	scores = map[string]int{}    // {}为定义空的映射
	scores = map[string]int{"李浪":100,"王健林":10,"马云":60}
	fmt.Println(scores)

	// 2、操作 -- 增、删、改、查
	fmt.Println(scores["李浪"])		// 输出为100，通过key找到了对应的value
	fmt.Println(scores["刘强东"])		// 如果访问不存在的key，则输出0，返回value对应的0值
		// 改
	scores["刘强东"] = 70
	fmt.Println(scores)
	scores["马化腾"] = 100			// 如果修改的时候，key不存在则添加
	fmt.Println(scores)

		// 删除 -> delete函数
	delete(scores,"马化腾")
	fmt.Println(scores)
	fmt.Println(len(scores))

	// 3、遍历 - 遍历是无序的
	for k, v := range scores {
		fmt.Println(k, v)
	}

	// 4、判断key是否存在
	if v, ok := scores["李浪"]; ok {
		fmt.Println(v)		// v代表value，如果变量ok为true，则返回value
	}

	// 5、key映射的value也可以是个映射。且key不能用切片，value可以
	var users map[string]map[string]string		// key为map[string]string，其也是一个映射
	users = map[string]map[string]string{"告诉你": {"老子":"天下","第一":"无敌"}}
	fmt.Printf("%T,%#v\n",users,users)
	fmt.Println(users["武大郎"] == nil)

	_, yes := users["武大郎"]
	fmt.Println(yes)
	users["武大郎"] = map[string]string{"来自":"北京"}
	fmt.Println(users)

	// 6、作业 - 统计投票数
	tongji := []string{"李浪","李浪","马化腾","马云","李浪","马化腾"}
	stat := make(map[string]int)
	for _, user := range tongji {		// 先遍历切片中的值给到变量user，_代表忽略索引
		if _, ok := stat[user]; !ok {		// 再把user值给到空映射stat，判断映射中值在不在
			stat[user] = 1
		} else {
			stat[user]++
		}
	}
	fmt.Println(stat)

	//简化
		/*for _,user := range tongji {
			stat[user]++
		}
		fmt.Println(stat)*/


}

