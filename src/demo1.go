package main

import (
	"B"
	"fmt"
	"reflect"
	"show"
)

const alen  =len(B.ZY)

const (
	aa=iota*2
	bb=iota*3
	cc
	ee
)
var a int
var(
	c string="张言"
	d int

)
func init()  {
	fmt.Println("主函数初始化")
}
func main() {
	fmt.Println("进入main函数")
	show.Show2()
	//test1.Test1()
	//var i int=1
	//var j float32=1.2
	//fmt.Println(+unsafe.Sizeof(i))
	//fmt.Println(+unsafe.Sizeof(j))

	var i int32
	var j float32
	var b bool
	var g complex64
	fmt.Print("int32:")
	fmt.Println(i)
	fmt.Print("float32")
	fmt.Println(j)
	fmt.Println(b)
	fmt.Println(d)
	a:=123

	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(g)
	fmt.Println(reflect.TypeOf(a))

	//var h int =3
	var l float32=3.01
	x:= int32(l)
	fmt.Println(reflect.TypeOf(x))

	fmt.Println(x)
	fmt.Println(B.Car)
	fmt.Println(B.ZAHANGYAN)
	fmt.Println(B.ZY)
	fmt.Println(B.Cat)
	fmt.Println(alen)

	fmt.Println(aa,bb,cc,ee)

	aaa:=[]string{"香蕉","苹果","小狗"}
	for k, v := range aaa {
		fmt.Println(k,v)
	}




}
