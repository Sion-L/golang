package main

import (
	"fmt"
	"math/rand"
	"time"
)

var input int    // 一般来说如果代码在块级别，则变量放在块里面，如果是重复的则放到外面来，可以减少重复申请内存的过程

func main() {
	// 设置随机数种子
	rand.Seed(time.Now().Unix())
	// 生成随机数0-100
	guess := rand.Intn(100)
	const maxGuessNum = 5
	//var isok bool

	for {    // 把所有代码放到一个for{}，游戏就会陷入死循环
		for i := 0; i < maxGuessNum; i++ {
			fmt.Print("请输入一个0-100的数字")
			fmt.Scan(&input)

			if guess == input {
				fmt.Printf("你猜对了,第%d次就对了\n",i+1)
				//isok = true
				break
			} else if guess > input {
				fmt.Printf("你猜小了,还有%d次机会\n",maxGuessNum-i-1)
			} else if guess < input {
				fmt.Printf("你猜大了,还有%d次机会\n",maxGuessNum-i-1)
			}
			//当最后一次也没才对，则输出
			if i == maxGuessNum-1 {
				fmt.Println("太垃圾了,游戏结束!")
			}

		}
		var text string       // 后面加个if选择语句，用来跳出循环
		fmt.Println("游戏是否继续?(y/n)")
		fmt.Scan(&text)
		if text != "y" {
			break
		}

	}



}