package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/**
 * @Author: zhangyan
 * @Date: 2019/5/5 23:31
 * @Version 1.0
 */
//1,1,2,3,5,8,13
//返回的实现了reader接口
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//函数的类型，但是只要是类型就能实现接口
type intGen func() int

//将函数作为接收者，实现Read接口
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	//TODO: 需要进行缓存，解决这个问题
	return strings.NewReader(s).Read(p)

}

//需要传入的是Reader类型的，需要实现read接口
func PrintFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	PrintFileContents(f)
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())

}
