// 遍历一段文字中每个字符出现的次数 rune => int
package main

import "fmt"

func main() {
	word := `sjkdjaksjdlkasljdkasnckjgasiudhasldwodiiojdiemfjnakdbyahuwidkjnmasklnckamojdiohqwertyuioasdfghjzxcvbn`
	word_stats := map[rune]int{}
	for _, ch := range word {
		if ch >= 'a' && ch <= 'z' {
			word_stats[ch]++
		}
	}

	countstats := map[int][]rune{}
	for ch, count := range  word_stats {
		/*if _,ok := countstats[count]; ok {		// 判断遍历出的值在不在
			countstats[count] = append(countstats[count],ch)
		} else {
			countstats[count] = []rune{ch}
		}*/
		// 以上代码可以简化 - 遍历在用key访问的时候返回的是0值，值为nil，nil切片也可以用append
		countstats[count] = append(countstats[count],ch)

	}
	fmt.Println(countstats)
}
