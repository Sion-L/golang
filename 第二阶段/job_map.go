/* 统计每个英文字符出现的次数
	每个字符在计算机里面都有一个编码的，英文字符的编码是ascii，A-Z，a-z
 */
package main

import "fmt"

func main() {
	article := `qwertyuioasdfghjklzxcvbnmrfghsjdtyhjsdfghsndcrtghdjffsghdjzxcvbnmqwertyutgbsxyhjnmwedfvujkdfgrtghnsmddfghdjktywuiekrfjnbvqwertyuioplqwqsadcsxvewsrdfgvsdfgbfsgdchfntgdhnfujdfmvfwsgdbnfmv,tedhfjydjfktgdhnf !dghjcfydhjfvmsgdhcfnvjsgdhcn`
	article_stat := map[rune]int{}
	for _, ch := range article {
		if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
			article_stat[ch]++
		}
	}
	fmt.Println(article_stat)
	for ch, cnt := range article_stat {
		fmt.Printf("%c:%d\n",ch,cnt)
	}

}
