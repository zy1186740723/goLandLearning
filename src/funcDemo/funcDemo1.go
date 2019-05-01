package main

import "fmt"

func main() {
	c:=3
	fmt.Println(c)
	C(&c)
	s1:=[]int{1,2,3,4,5}
	a,b:=1,2
	A(a,b,3,4,5)
	B(s1)
	fmt.Println(a,b)
	fmt.Println(s1)
	//函数作为变量,函数类型的使用
	d:=D
	d()

	//代码块变为函数类型，赋值给了e
	e:= func() {
		fmt.Println("匿名函数func e")
	}
	e()

	//闭包函数
	f:=closeure(10)
	fmt.Println(f(1))

}
//1、不定长变参 2、值拷贝

func A(a ...int) {
	a[0]=3
	a[1]=4
	fmt.Println(a)
}

//内存地址的传递
func B(s []int)  {
	s[0]=5
	s[1]=6
	fmt.Println(s)
}
func C(a *int)  {
	*a=2
	fmt.Println(*a)

}
func D()  {
	fmt.Println("FUNC D")
}
//闭包函数
func closeure(x int) (func(int) int){
	//返回一个匿名函数
	fmt.Printf("%p\n",&x)
	return func(y int) int {
		fmt.Printf("%p\n",&x)
		return x+y
	}
}
