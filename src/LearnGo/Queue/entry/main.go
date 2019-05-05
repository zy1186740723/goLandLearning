package main

import (
	queue "LearnGo/Queue"
	"fmt"
)

/**
 * @Author: zhangyan
 * @Date: 2019/5/5 10:51
 * @Version 1.0
 */
func main() {
	q := queue.Queue{1}

	//改变了q中的值 ，与开始的不是一个q
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	q.Push("abc")
	fmt.Println(q.Pop())

}
