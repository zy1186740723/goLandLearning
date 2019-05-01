package main

import "fmt"

type A1 struct {
	B1 //匿名字段
	Name string
}

type B1 struct {
	Name string
}

type C1 struct {
	Name string
}
func main() {
	a:=A1{Name:"ZHANG",B1:B1{Name:"yan"}}
	//如果匿名字段与外层结构有同名字段，应该如何操作
	fmt.Println(a.Name,a.B1.Name)
}
