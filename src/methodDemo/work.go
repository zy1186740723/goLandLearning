package main

import "fmt"

type TZ1 int

func main() {
	var a TZ1
	a = 0
	a.Increase()
	fmt.Println(a)
	a.Increase2(200)
	fmt.Println(a)

}

func (a *TZ1) Increase() {
	for i := 0; i < 100; i++ {
		*a += 1
	}
}

func (a *TZ1) Increase2(num int) {
	//类型变换
	// 以int为底层，增加丰富的功能
	*a += TZ1(num)

}
