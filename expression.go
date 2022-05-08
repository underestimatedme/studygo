package main

import "fmt"

/**
 * go的25个关键字如下:
 * break,default,func,interface,select,case,defer,go,map,struct,chan,else,goto,package,switch,const,fallthrough,if,range,type,continue,for,import,return,var
 * 其中与java有差异的func方法/select/defer/go/map/struct/chan/const/fallthrough/range/var
 */
func defIncrement() {
	a := 1
	//++a  // 前置的方式go取消了
	a++
	p := &a // 将a的地址赋给p
	*p++    // 类似于c的指针操作语法
	fmt.Println(a)
	/**
	 * & 取地址运算符 - 用于获取对象地址
	 * * 指针运算符 - 间接引用目标对象
	 * **T 二级指针
	 */
}

func defInit() {
	// 定义一个结构体数据类型
	type student struct {
		name   string
		height int
		age    int
	}
	// 使用花挂号来初始化结构体
	var joey student = student{"joey", 120, 5}
	judy := student{"judy", 115, 5}
	fmt.Println(joey)
	fmt.Println(judy)
}

func main() {
	defIncrement()
	defInit()
}
