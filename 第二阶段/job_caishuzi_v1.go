package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置随机数种子
	rand.Seed(time.Now().Unix())
	// 生成随机数0-100
	guess := rand.Intn(100)
	const maxGuessNum = 5
	//var isok bool

	// 定义猜的次数，i<=maxGuessNum i要小于最大可猜次数5
	for i := 0; i < maxGuessNum; i++ {
		var input int
		fmt.Print("请输入一个0-100的数字")
		fmt.Scan(&input)			// 定义输入猜的值

	if guess == input {				// if选择，对比猜的值和随机值的大小
		fmt.Printf("你猜对了,第%d次就对了\n",i+1)
		//isok = true
		break			// 如果猜对了就退出循环
	} else if guess > input {
		fmt.Printf("你猜小了,还有%d次机会\n",maxGuessNum-i-1)
	} else if guess < input {
		fmt.Printf("你猜大了,还有%d次机会\n",maxGuessNum-i-1)
	}
	//当最后一次也没才对，则输出
	if i == maxGuessNum-1 {
		fmt.Println("太垃圾了,游戏结束!")
	}

	// 或者用添加一个布尔变量isok
	/*if !isok  {
		fmt.Println("  太垃圾了,游戏结束")
	}*/
	}

}