package main

import "fmt"

func main() {
	for i:=0;i<3 ;i++  {
		defer fmt.Println(i)
		defer func() {
			fmt.Println(i)
		}()
	}
	defer fmt.Println("b")
	defer fmt.Println("c")

}

//go语言中的错误处理机制
//panic/recover来处理错误 recover必须是defer调用的函数有效
