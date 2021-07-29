package main

import "fmt"

func main() {
	A := 2
	B := A
	B = 3
	fmt.Println(A,B)

	// 指针
	var cc *int = &A    // 跟下面的含义一样,*int - int类型指针，*string - string类型指针....
	C := &A     // 把A的内存地址取出来赋值给变量C,&可以获取指针的地址
	fmt.Printf("%T %T %p\n",C,cc,cc)   // %p为打印指针，%v可打印所有数据类型的指针

	fmt.Println(*cc)    // *可以取指针的值,既内存地址中的值
	*cc = 4  		// 修改指针的值，对应A的值也会改变，即通过修改指针的值去修改变量的值
	fmt.Println(A)

	//test
	E := 3
	F := E
	F = 2
	fmt.Println(E,F)
	dd := &E  	// dd的值为内存地址的值，即跟E共用一个内存地址
	*dd = 4			// 此时修改dd内存地址的值，则E的值也发生改变
	fmt.Println(A)
}
