package main

import "fmt"

func main()  {
	var fs=[4]func(){}

	for i:=0;i<4 ;i++  {
		//这里的i是值拷贝
		defer fmt.Println("defer i=",i)
		//闭包的思想
		defer func() {fmt.Println("defer_closure i=",i)}()
		//i从外层读入，是引用地址，闭包的思想
		fs[i]= func() {
			fmt.Println("closure i=",i)
		}
	}
	for _,f := range fs {
		f()
	}
	}



