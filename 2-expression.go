package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

/**
 * go的25个关键字如下:
 * break,default,func,interface,select,case,defer,go,map,struct,chan,else,goto,package,switch,const,fallthrough,if,range,type,continue,for,import,return,var
 * 其中与java有差异的func方法/select(信道)/fallthrough(控制语句) map/struct/const/range/var(变量和数据) defer(延迟)/go(协程)/chan()
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

func defArr() {
	// array定义 []int{} 这种方式
	c := []int{1, 2, 3}
	d := []int{1, 2, 3, 4, 5}
	fmt.Println(c, d)
}

func defIf() {
	// if-else if-else 流程和Java一样，只是可省略括号
	var age int = 6
	if age < 7 {
		fmt.Println("学龄前")
	} else if 7 <= age && age < 13 {
		fmt.Println("小学生")
	} else {
		fmt.Println("其他")
	}
	if age++; age >= 7 {
		fmt.Println("小学生及以上")
	}
	// go语言不支持以下的三元运算符
	//a := (age>=7 ? "":"")
	//省略: go的switch和其他语言的差异在于只要case执行成功就不会执行后面的语句,只有case全失败才会执行default，更符合普通使用的判断场景[备注:fallthrough强制进入下一个判断流程]，不过我推测应该使用的很少，省略这部分
}

func defFor() {
	// go中只有for循环，没有while和do-while.
	for i := 0; i < 3; i++ { // 普通的for循环
		fmt.Println(i)
	}
	x := 0
	for x < 10 { // 类似于while(x<10){}
		x++
	}
	for { // while(true){}
		break
	}
	// for range
	data := []string{"a", "b", "c"}
	for i, s := range data {
		fmt.Println(i, s)
		if i == 0 {
			data[0] = "b"
		}
	}

	// range会复制目标数据，循环时数组发生变化变量不变
	m := [3]int{10, 20, 30}
	for i, x := range m {
		if i == 0 {
			m[0] += 100
			m[1] += 200
			m[2] += 300
		}
		//x: 10,data: 110
		//x: 20,data: 220
		//x: 30,data: 330
		fmt.Printf("x: %d,data: %d\n", x, m[i])
	}

}

// func不支持嵌套
func add(x, y int) int {
	return x + y
}

func exec(f func(x, y int) int, x int, y int) int {
	return f(x, y)
}

// 指针是go与其他语言最显著的区别
func testPointer() *int {
	a := 10
	return &a
}

// 匿名函数, 和js基本一样
func testAnonymous() {
	func(s string) {
		fmt.Println(s)
	}("test it")
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))
}

// 闭包closure，defer todo:暂时省略

// 错误处理 error handle
// 此处的error为官方接口
func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("divide by zero")
	}
	return x / y, nil
}

// panic-recover-类似于try...catch...，但是感觉比try catch难用很多
func testPanic() {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()
	panic("holy shit")
	// 下面这句执行不到了
	fmt.Println("test shit")
}

func main() {
	defIncrement()
	defInit()
	defArr()
	defIf()
	defFor()
	fmt.Println(exec(add, 1, 2))
	testAnonymous()
	res, err := div(1, 0)
	fmt.Println(res, err)
	testPanic()
}
