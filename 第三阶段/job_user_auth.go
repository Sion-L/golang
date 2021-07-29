package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	maxAuth = 3   // 设置最多认证三次
	passwd = "123.com"
)

func inputString(prompt string) string {		// 提供输入功能的函数，要使用可直接调用
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return strings.TrimSpace(input)		// 把字符串前后空格去掉
}

func inputUser() map[string]string {			// 定义添加和修改用户信息的函数
	user := map[string]string{}			// 用映射存储信息
	user["name"] =  inputString("请输入名字：")
	user["age"] = inputString("请输入年龄：")
	user["tel"] = inputString("请输入你的电话号码：")
	user["addr"] = inputString("请输入你的地址：")
	return user
}

func printUser(out int,user map[string]string) {		// 定义输出用户信息的函数
	fmt.Println("ID:",out)
	fmt.Println("名字:",user["name"])
	fmt.Println("年龄:",user["age"])
	fmt.Println("电话:",user["tel"])
	fmt.Println("地址:",user["addr"])
}

// 从命令行输入密码并进行验证
func auth() bool {  // 通过bool类型的返回值去判定认证结果
	var input string
	for i:= 0; i <= maxAuth; i++ {
		fmt.Print("请输入的你的密码：")
		fmt.Scan(&input)
		if passwd == input {
			return true
		} else {
			fmt.Println("密码错误")
		}
	}
	return true
}

// 查询功能
func query(users map[int]map[string]string) {
	q := inputString("请输入查询的内容：")		// 直接调用提供输入功能的函数
	title := fmt.Sprintf("%5s|%20s|%20s|%20s|%20s","ID","Name","Age","Tel","Addr")  // 结构占位
	fmt.Println(title)
	fmt.Println(strings.Repeat("-",len(title)))
	for k, v := range users {
	   // 判断users的value是否包含查询的内容
		if strings.Contains(v["name"],q) || strings.Contains(v["tel"],q) || strings.Contains(v["age"],q) || strings.Contains(v["addr"],q) {						// 判断字符串是否包含指定的值
			fmt.Printf("%#v%5s|%20s|%20s|%20s|%20s",k,v["id"],v["name"],v["age"],v["tel"],v["addr"])
			fmt.Println()
	   }

}
}

// 添加功能
func getId(users map[int]map[string]string) int { // 获取users中最大的id
	var id int
	for k,_ := range users {
		if id < k {
			id = k
		}
	}
	return id + 1
}

func addn(users map[int]map[string]string) {
	id := getId(users)
	user := inputUser()
	users[id] = user
	fmt.Println("添加成功")
}


// 修改功能
func modify(users map[int]map[string]string) {
	if id, err := strconv.Atoi(inputString("请输入修改用户id：")); err == nil {	 // 字符串转化为int并判断有没有错误
		if user, ok := users[id]; ok {			// 判断用户id存不存在
			fmt.Println("你将修改的用户信息:")
			printUser(id,user)
			input := inputString("确定修改吗?(Y/N)")
			if input == "Y" || input == "y" {
				user := inputUser()
				users[id] = user
			}
		} else {
			fmt.Println("用户id不存在")
		}

	} else {			// 判断id正不正确
		fmt.Println("输入的id不正确")
	}
}

// 删除功能
func del(users map[int]map[string]string) {
	if id, err := strconv.Atoi(inputString("请输入删除的用户id：")); err == nil { // 字符串转化为int并判断有没有错误
		if user, ok := users[id]; ok {			// 判断用户id存不存在
			fmt.Println("你将删除的用户信息:")
			fmt.Println(user,id)
			input := inputString("确定删除吗?(Y/N)")
			if input == "Y" || input == "y" {
				delete(users,id)
				fmt.Println("删除成功")
			}
		} else {
			fmt.Println("用户id不存在")
		}

	} else {			// 判断id正不正确
		fmt.Println("输入的id不正确")
	}
}


func main() {
	if !auth() {
		fmt.Println("密码输入正确，程序退出")
		return
	}
	// 定义菜单变量
	menu := `***********************************************************************			
1.查询
2.添加
3.修改
4.删除
5.退出
***************************************************************************`


	// 映射users存放id、name、age、tel、addr...
		//users := map[int][]string{}        也可用切片
		//users := map[int][5]string{}				int为id号，[5]长度为5的数组存放地址、电话、名字、生日
	users := map[int]map[string]string{}			//地址、电话、名字、生日也可用映射来存放,映射可直接用delete函数删除元素

	callback := map[string]func(map[int]map[string]string) {
		"1" : query,
		"2" : addn,
		"3" : modify,
		"4" : del,
		"5" : func(users map[int]map[string]string) {
			os.Exit(0)
		},
	}
	fmt.Println("欢迎进入用户管理系统")
//END:
	for {
		fmt.Println(menu)
		if callback, ok :=  callback[inputString("请输入指令:")]; ok {
			callback(users)
		} else {
			fmt.Println("指令错误")
		}


		/*switch inputString("请输入指令:") {		// 直接调用提供输入功能的函数
		case "1":
			query(users)
		case "2":
			addn(users)
		case "3":
			modify(users)
		case "4":
			del(users)
		case "5":
			break END
		default:
			fmt.Println("指令错误")
		}
*/
		}
}