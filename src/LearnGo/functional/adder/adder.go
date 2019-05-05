package main

import (
	"fmt"
)

/**函数式编程1、闭包 2、函数实现接口 3、使用函数遍历二叉树
 * @Author: zhangyan
 * @Date: 2019/5/5 23:08
 * @Version 1.0
 */
//闭包的概念
func adder() func(int) int {
	//sum是自由变量，v是局部变量
	//自由变量会连到结构中的变量中去，形成一个闭包
	//函数返回的是一个闭包
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (i int, i2 iAdder) {
		return base + v, adder2(base + v)
	}

}
func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}
	b := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, b = b(i)
		fmt.Printf("%d %d", i, s)
		fmt.Println()
	}
}
