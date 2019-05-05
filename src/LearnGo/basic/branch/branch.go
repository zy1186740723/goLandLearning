package main

import (
	"fmt"
	"io/ioutil"
)

//switch的用法
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprint("Wrong score:%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}
func main() {
	//读取文件内容
	const filename = "abc.txt"
	//contents,err:=ioutil.ReadFile(filename)
	//特殊之处 if的条件里可以赋值多个语句
	//作用域在if语句内
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	fmt.Println(
		grade(0),
		grade(10),
		grade(90),
		grade(100))

}
