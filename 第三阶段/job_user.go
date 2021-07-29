/*
用户信息存储
=>内存
=>结构 map => []
=>用户 id name age tel addr
	  map
	  值类型使用 string

用户的修改
请输入要修改的用户名
users[id] => 用户是否存在,不在提示用户不正确
打印用户信息，提示用户是否确认修改（Y/N）
Y提示用户输入修改后内容

用户的删除
请输入需要删除的用户id
users[id] => 用户是否存在,不在提示用户不正确
打印用户信息，提示用户是否确认删除（Y/N）
delete()
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 添加用户
func zeng(pk int,users map[string]map[string]string) {
	var (
		id string = strconv.Itoa(pk)	// 字符串转int类型
		name string
		age string
		tel string
		addr string
	)
	fmt.Println(id)

	fmt.Print("请输入姓名：")
	fmt.Scan(&name)

	fmt.Print("请输入你的年龄：")
	fmt.Scan(&age)

	fmt.Print("请输入你的电话号码：")
	fmt.Scan(&tel)

	fmt.Print("请输入你的地址：")
	fmt.Scan(&addr)

	users[id] = map[string]string {
		"id" : id,
		"name" : name,
		"tel" : tel,
		"age" : age,
		"addr" : addr,
	}
	fmt.Println(id,name,age,tel,addr)

}

// 查询用户,q为空则查询全部.非空则输出name、tel、addr中包含q内容的显示
func query(users map[string]map[string]string) {
	var q string
	fmt.Println("请输入查询信息")
	fmt.Scan(&q)
	title := fmt.Sprintf("%5s|%20s|%20s|%20s|%50s","ID","Name","Age","Tel","Addr")  // 结构占位
	fmt.Println(title)			// 打印结构
	fmt.Println(strings.Repeat("-",len(title)))		// -- 分隔
	for _, user := range users {
		if q == "" || strings.Contains(user["name"],q) || strings.Contains(user["age"],q) || strings.Contains(user["tel"],q) || strings.Contains(user["addr"],q) {
			fmt.Printf("%5s|%20s|%20s|%20s|%50s",user["id"],user["name"],user["age"],user["tel"],user["addr"])
			fmt.Println()
		}
		}
}

func main() {
	// 存储用户信息
	users := make(map[string]map[string]string)
	id := 0
	fmt.Println("用户系统")
	for {
		var op string
		fmt.Print(`
1.创建用户
2.修改用户
3.删除用户
4.查询用户
5.退出
请输入指令：`)

		fmt.Scan(&op)
		if op == "1" {
			id++			// 定义自增id，创建用户的时候自增
			zeng(id,users)
		} else if op == "2" {

		} else if op == "3" {

		} else if op == "4" {
			query(users)
		} else if op == "5" {
			break
		} else {
			fmt.Println("指令错误")
		}

	}
}
