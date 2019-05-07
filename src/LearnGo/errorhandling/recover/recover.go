package main

import (
	"fmt"
)

/**
 * @Author: zhangyan
 * @Date: 2019/5/6 19:55
 * @Version 1.0
 */

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occurred:", err)
		} else {
			panic(r)
		}
	}()
	//panic(errors.New("this is an error"))
	a := 0
	b := 5 / a
	fmt.Println(b)

}
func main() {
	tryRecover()
}
