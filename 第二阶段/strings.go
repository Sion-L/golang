package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1、strings包中的函数
	fmt.Println(strings.Compare("abc","bac"))			// abc < bac ，所以返回值为-1
	fmt.Println(strings.Contains("abc","bc"))      // abc是否包含bc，包含则返回true,不包含则返回false
	fmt.Println(strings.ContainsAny("abc","ad"))    // ad包含abc中的a或任意字符，则返回true
	fmt.Println(strings.Count("aaavcasdada","a"))   // 统计a出现的次数
	fmt.Printf("%q\n",strings.Fields("abc\nds\reee\tdsd\f"))	//按空白字符分割（空格,\n,\r,\f,\t）
	fmt.Println(strings.HasPrefix("abcd","ab"))		// 如果abcd是以ab开头的，则返回true
	fmt.Println(strings.HasSuffix("abcd","cd"))	// 如果abcd以cd结尾则返回true
	fmt.Println(strings.Index("abcda","cda"))		// 返回子串cda在字符串abcda中第一次出现的位置,如果找不到则返回-1；如果str为空，则返回0
	fmt.Println(strings.LastIndex("abcde","de"))

	fmt.Println(strings.Split("ab;c;def",";"))		// 把;换成空格分割
	fmt.Println(strings.Join([]string{"def","ard","yui"},":"))	// 将字符串连接起来，用:连接

	fmt.Println(strings.Repeat("abc",3))		// 将abc重复三次
	fmt.Println(strings.Replace("abcbacabc","ab","xxx",1))		// 把abcbacabc中的ab替换成xxx，并只替换第一个

	fmt.Println(strings.ToLower("abCd"))   // 全部的字符转为小写
	fmt.Println(strings.ToUpper("abced"))	// 全部转为大写
	fmt.Println(strings.Title("abc, efg"))	// 首字母大写

	fmt.Println(strings.Trim("abjdabsdba","ab"))	// 去掉左右两边两个含有ab任意一个单词的字符
	fmt.Println(strings.TrimSpace(" abcd xxx \n \r"))   // 把左右两边的空白字符和\r,\t,\n去掉




}
