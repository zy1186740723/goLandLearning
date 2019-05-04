package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//两个返回值常用的场景，用于返回error
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil

	default:
		return 0, fmt.Errorf("Unsupported op:%s", op)
	}

}

//带余数的触发
func div(a, b int) (q, r int) {
	return a / b, a % b
}

//函数式编程的方法，参数也可以是一个函数
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name() //反射获取函数的名称
	fmt.Printf("calling func %s with args"+"(%d,%d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))

}

//可变参数列表
//求可变参数列表的和
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s

}
func swap(a, b int) (int, int) {
	return b, a
}
func swap2(a, b *int) {
	*b, *a = *a, *b

}
func main() {
	fmt.Println(eval(3, 4, "*"))
	fmt.Println(div(13, 3))
	q, r := div(15, 4)
	fmt.Println(q, r)
	fmt.Println(eval(3, 4, "X"))
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(apply(pow, 3, 4))
	//匿名函数的方法
	fmt.Println(apply(
		func(a, b int) int {
			return a + b
		}, 3, 4))
	fmt.Println(sum(1, 23, 4, 5, 7))

	a, b := swap(3, 4)
	fmt.Println(a, b)
	c, d := 5, 6
	swap2(&c, &d)
	fmt.Println(c, d)

}
