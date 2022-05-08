package main

// 引用其他包使用import ""方式
import (
	"fmt"
	"os"
)

/**
 * go语言的基础语法，go的语法很灵活
 * 和java相同基础类型有bool, byte, int8, int16, int32, rune, float32, float64,
 * 除此以外还有int,int64,complex,uintptr(存储指针),string(初始值为""),array,struct(结构体，和c一脉相承),function,interface等类型
 * 引用类型包括: map,slice,channel
 *
 */
var x = "xx" // 全局变量
func defVar() {
	// 定义变量有2种方式
	var a = "123"
	var a0 int // go语句结尾不用写封号
	b, c := "4", "5"
	// go语言中不允许定义了变量而不使用,否则编译不通过
	fmt.Println(a, a0, b, c)
	// 常量可以定义而不使用
	const d = "cc"

	println(&x, x)
	x := "yy"
	println(&x, x)

	{
		// 不同作用域可重新定义变量
		b, c := "6", "7"
		fmt.Println(b, c)
	}

	// 对err进行退化赋值处理
	f, err := os.Open("/opt/file/test.txt")
	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	println(err, n)
}

func defReference() {
	// 数组定义(第一次看到时语法有点不太能接受)
	s := []int{1, 2, 3, 4}
	s = append(s, 5)
	// 切片语法和python完全一样
	fmt.Println(s[1:])
	// map的语法
	m := make(map[string]int)
	m["a"] = 1
	// 判断map是否包含某个元素
	if _, ok := m["b"]; ok {
		// do nothing
	} else {
		m["b"] = 2
	}
	fmt.Println(m)
}

/**
 * go语言中为啥要保留struct这个结构体呢，有点不太明白，数据结构java中一般都是定义个class解决的
 *
 */
func defType() {
	type (
		user struct {
			name string
			age  uint8
		}
		event func(string) bool
	)

	u := user{"tom", 20}
	fmt.Println(u)

	var f event = func(s string) bool {
		println(s)
		return s != ""
	}

	f("abc")
}

func main() {
	//defVar()
	defReference()
	defType()
}
