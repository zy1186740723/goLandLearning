package main

import "fmt"
//类型别名
type TZ int //对int有一些高级的操作的时候，可以把int作为底层
type A struct {
	Name string
}

type B struct {
	Name string

}
type C struct {
	name string //造package中的公有的，超出这个包就必须用大写的公有字段了

}
func main() {
	a:=A{Name:"zhangyan"}
	a.print()
	fmt.Println(a.Name)

	b:=B{}
	b.print()
	fmt.Println(b)

	var c TZ
	c.print2()
	//第二种调用方法
	(*TZ).print2(&c)
}
//go语言无法重载
func (a *A) print()  {
	a.Name="aa"
	fmt.Println(a)
}
func (b B) print()  {
	b.Name="bb"
	fmt.Println(b)
}

//Reciver func的第一个参数
func (a *TZ) print2() {
	fmt.Println(*a+1)

}

func (c *C) print() {
	//方法具有高的访问权限，可访问中的私有成员
	c.name="123"
	fmt.Println(c.name)

}
//方法与函数的互通