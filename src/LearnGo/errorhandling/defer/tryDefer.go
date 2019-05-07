package main

import (
	"LearnGo/functional/fib"
	"bufio"
	"fmt"
	"os"
)

/**
 * @Author: zhangyan
 * @Date: 2019/5/6 16:11
 * @Version 1.0
 */

func tryDefer() {

	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		//退出的时候i是等于30
		if i == 30 {
			panic("printed too many")
		}

	}
	//defer中相当于有个栈 是先进后出的
	//有了defer不怕return和panic的出现，语句仍然被执行
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error")
	fmt.Println(4)

}

/*defer确保在函数结束时发生，参数在defer语句时计算
defer列表为后进先出
*/
func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//写入内存中，当内存大到一定程序全部倒过去
	writer := bufio.NewWriter(file)
	//从内存写入文件中去
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		//Fp 可传入文件参数
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fib.txt")
	tryDefer()
}
