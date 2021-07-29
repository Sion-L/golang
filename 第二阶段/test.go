package main

import (
	"fmt"
)

func main() {
	word := `sdjaksjdiwadjnsadxkjasrfldsjdkajskdjiwajcsjnznddjsajwjekjskandnxanskedjsdnasndnajwebassjxnajkejkwjkd`
	word_stat := map[rune]int{}
	for _,ch := range word {
		if ch >= 'a' && ch <= 'z' {
			word_stat[ch]++
		}
	}

	countstat := map[int][]rune{}
	for ch, count := range word_stat {
		countstat[count] = append(countstat[count],ch)
	}
	fmt.Println(countstat)


}
	// 总结：映射的含义、映射的定义、映射的字面量（赋值）、映射的操作（增删改查）、遍历映射、判断key是否存在
