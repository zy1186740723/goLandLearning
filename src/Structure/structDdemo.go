package main

import "fmt"


type person struct {
	Name string
	Age int
}
func main(){
	a:=person{}
	a.Age=10
	a.Name="zhang"

	//结构的初始化，习惯性用取地址的方法
	b:=&person{
		Name:"liu",
		Age: 12,
	}
	fmt.Println(a)
	//值的拷贝
	A(a)
	fmt.Println("值拷贝后的结果",a)
	fmt.Println(b)
	//地址传递
	B(b)
	fmt.Println("地址传递后的结果",b)

	//匿名结构
	c:= &struct {
		Name string
		Age int
	}{
		Name:"c",
		Age:12,
	}

	fmt.Println(c)

}

func A(per person)  {
	per.Age=13
	fmt.Println("值拷贝调用函数中",per)
}

func B(per *person)  {
	per.Age=13
	fmt.Println("指针传递",per)

}

func C(per *person)  {
	per.Age=15
	fmt.Println("指针传递",per)
}
