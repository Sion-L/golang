package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成随机数的模块
// 时间模块
func main() {
	rand.Seed(time.Now().Unix())   // 随机数的种子
	fmt.Println(rand.Int())     // 输出生成的随机数
 	fmt.Println(rand.Int() % 100)   // 取余，生成0-100的随机数
 	fmt.Println(rand.Intn(100))   // 限制随机数的上限为100
}