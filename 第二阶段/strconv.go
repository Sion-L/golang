// 此包是做字符串转化的
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串到其它类型的转换
		// 1、到bool类型的转换,只能转化为true或者false
	if v, err := strconv.ParseBool("true"); err == nil {
		fmt.Println(v, err) // 判断err是否为nil，也就是是否有错误信息，是nil的话则打印v
		} else {
			fmt.Println(err)	// 有错误信息则输出错误信息
	}

		// 2、字符串到int类型的转换
	if v, err := strconv.Atoi("1023"); err == nil {
		fmt.Println(v)
	}	else {
		fmt.Println(err)
	}
			// 或者使用ParseInt，这个函数需要指定进制类型和字节数
	if v, err := strconv.ParseInt("64",16,64); err == nil {
		fmt.Println(v)		// 16*6+4 = 100
	}	else {
		fmt.Println(err)
	}

		// 3、到float的转换
	if v, err := strconv.ParseFloat("1.1",64); err == nil {
		fmt.Println(v)
	}


	// fmt.Sprintf也可以转换
	sd := fmt.Sprintf("%d",12)	// 转int
	fmt.Println(sd)

	sf := fmt.Sprintf("%.3f",12.02)  // 转float,.3为保留三位小数
	fmt.Println(sf)


	// 其它类型转为字符串
	fmt.Printf("%q\n",strconv.FormatBool(true))
	fmt.Printf("%q\n",strconv.Itoa(12))
	fmt.Printf("%q\n",strconv.FormatInt(12,16))
	fmt.Printf("%q\n",strconv.FormatFloat(10.1,'E',-1,64))  // 科学计数法




}

// 总结：strconv.ParseFloat 把