package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//函数外的定义，包内部的变量
var aa = 3
var ss = "kkk"

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c = 2, true, "def"
	fmt.Println(a, b, c)
}

func variableShorter() {
	a, b, c, d := 3, 4, true, "def"
	fmt.Println(a, b, c, d)

}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)

}

//强制类型转换
func triangle() {
	var a, b int = 3, 4
	var c int
	//进行两次强制类型转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)

}

//常量的定义问题
func Consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	//此处的ab就是文本 不需要进行强制类型转换
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

//枚举的问题
func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
	)
	const (
		java1 = iota
		java2
		java3
		_
		java4
	)
	//b,kb,mb,gb
	const (
		b = 1 << (10 * iota) //左移十个iota位
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python)
	fmt.Println(java1, java2, java3, java4)
	fmt.Println(b, kb, mb, gb, tb, pb)

}
func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	euler()
	triangle()
	Consts()
	enums()
}
