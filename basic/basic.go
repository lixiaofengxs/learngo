package main

import (
	"fmt"
	"math"
)

/*
内建变量类型：
1.bool string
2.(u)int8 (u)int16 (u)int32 (u)int64 uintptr
3.byte rune (char 32位)
4.float32 float64 complex64 complex128
*/

/*
强制类型转换：go只有强制类型转换，没有隐式转换
*/

// 包内变量。go语言没有全局变量的概念，作用域只是包内
var aa = 3
var ss = "kkk"

// 常量也可以定义在函数外，在整个包中都可以使用
// go语言中常量名一般不会大写，因为大写在go语言是由特殊含义的
const filepath = "./file/"

// const数值可以作为各种类型使用
const (
	file = "aaa"
	m, n = 3, 4
)

// 一般这样写
var (
	a1 = 1
	a2 = "s"
	a3 = true
)

// 在函数外面定义变量不能使用简洁形式
// bb := true

func variableZeroValue() {
	// 使用var关键字定义变量，变量名在前，变量类型在后
	var a int
	var s string

	// go语言变量定义后有初始值。a-0, s-""
	fmt.Println(a, s)

	// %q会把空串的引号打印出来
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	// 变量一旦定义，必须要使用，否则报错
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	// 变量类型自动推导
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	// 变量定义更简单的写法，只能在函数内部这样定义
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) //必须强制转换
	fmt.Println("c = ", c)
}

func constVariable() {
	const filename = "abc.txt"
	// 定义常量可以指定类型，也可以不指定
	// a, b 类型不确定。go中的常量类似于文本替换
	const a, b = 3, 4
	const c, d int = 3, 4

	var c1 int
	//a, b没定义类型，因此既可以做int，也可以做float。就是一个文本
	c1 = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c1)
}

// go语言没有特殊的枚举关键字，一般就是使用const定义
func enumType() {
	// const一定要给它值
	// 普通枚举类型
	//const (
	//	cpp    = 0
	//	java   = 1
	//	python = 2
	//	golang = 3
	//)

	// 自增枚举类型
	const (
		cpp = iota //go语言提供的自增值
		_
		java
		python
		golang
	)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello World!")

	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()

	triangle()
	constVariable()
	enumType()

	fmt.Println(a1, a2, a3)
}
