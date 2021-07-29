// 汉诺塔游戏
package main

import (
	"fmt"
)

func test(a,b,c string, layer int) {
	if layer  == 1 {			// 第一层把圈圈从a柱移动到c柱
		fmt.Println(a, "->", c)
		return		// golang中没定义返回值则直接return
	}
	// a layer-1（第二层），借助c移动到b
	test(a,c,b, layer-1)
	fmt.Println(a,"->",c)
	test(b,a,c, layer-1)
}

func main() {
	fmt.Println("1层:")
	test("A","B","C",1)
	fmt.Println("2层:")
	test("A","B","C",2)
	fmt.Println("3层:")
	test("A","B","C",3)
	fmt.Println("4层:")
	test("A","B","C",4)

}

